package Service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
	"myGoEthereum/Helper/ConfigHelper"
	"myGoEthereum/Helper/LogHelper"
	"myGoEthereum/Model/BaseModel"
	"myGoEthereum/Model/CommonModel"
	"myGoEthereum/Model/GormModel"
	"myGoEthereum/Model/ResultCode"
	"myGoEthereum/Repository/Repository"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/golang-jwt/jwt"
)

func GetNewBlocks() (response CommonModel.ApiResponseWithData) {
	dbNumber := Repository.GetDbMaxBlockNumber()
	begin := dbNumber - ConfigHelper.GetInt64("UnstableN")
	if begin < 0 {
		begin = 0
	}
	end := getCurrentBlockNumber()
	InsertBlockFromRpc(uint64(begin), end)
	return
}

func InitialDbData() (response CommonModel.ApiResponseWithData) {
	// truncate table
	Repository.TruncateBlocksTransactionsReceiptLogs()

	// init data
	number := getCurrentBlockNumber()
	begin := number - uint64(ConfigHelper.GetInt("InitialBlockSize"))
	if begin < 0 {
		begin = 0
	}
	InsertBlockFromRpc(begin, number)
	return
}

func InsertBlockFromRpc(begin, end uint64) {
	blocks, transactions, receiptLogs := GetBlockFromRpcByNumberRanmge(begin, end) //1009984
	if len(blocks) > 0 {
		Repository.InsertOnConflictUpdate(blocks)
	}
	if len(transactions) > 0 {
		Repository.InsertOnConflictUpdate(transactions)
	}
	if len(receiptLogs) > 0 {
		Repository.InsertOnConflictUpdate(receiptLogs)
	}
	return
}

func getCurrentBlockNumber() uint64 {
	client, err := getEthclientClient(ConfigHelper.GetString("Endpoint"))
	if err != nil {
		panic(err)
	}
	number, err := client.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}
	return number
}

func GetBlockFromRpcByNumberRanmge(begin, end uint64) ([]GormModel.Block, []GormModel.Transaction, []GormModel.ReceiptLog) {
	concurrencyLimit := runtime.NumCPU() * ConfigHelper.GetInt("MultipleOfCpuCores")
	input := make(chan uint64, concurrencyLimit*2) //設size避免用太多記憶體
	var myBlock GormModel.BlocksTransactionsMutex

	go func(begin, end uint64) {
		defer close(input)
		for i := end; i > begin; i-- {
			input <- i
		}
	}(begin, end)

	var wg sync.WaitGroup
	wg.Add(concurrencyLimit)

	client, err := getEthclientClient(ConfigHelper.GetString("Endpoint"))
	if err != nil {
		panic(err)
	}

	for i := 0; i < concurrencyLimit; i++ {
		go func() {
			defer wg.Done()
			for v := range input {
				block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(v))
				if err != nil {
					fmt.Println("[Error]", err)
					continue
				}
				parseBlock := parseBlockHeader(block)
				parseTransactions := parseTransaction(block.Transactions(), &parseBlock.Number)

				var txHash []common.Hash
				var receiptLogs []GormModel.ReceiptLog
				linq.From(block.Transactions()).SelectT(func(v *types.Transaction) common.Hash { return v.Hash() }).ToSlice(&txHash)
				if txHash != nil && len(txHash) > 0 {
					receiptLogs = getTransactionsReceiptLogs(client, txHash, parseBlock.Number)
				}

				// lock
				myBlock.Mux.Lock()
				if parseBlock.Number > -1 {
					myBlock.Blocks = append(myBlock.Blocks, parseBlock)
				}
				if len(parseTransactions) > 0 {
					myBlock.Transactions = append(myBlock.Transactions, parseTransactions...)
				}
				if len(receiptLogs) > 0 {
					myBlock.ReceiptLogs = append(myBlock.ReceiptLogs, receiptLogs...)
				}
				myBlock.Mux.Unlock()
			}
		}()
	}
	wg.Wait()
	return myBlock.Blocks, myBlock.Transactions, myBlock.ReceiptLogs
}

func getEthclientClient(host string) (*ethclient.Client, error) {
	ctx, err := rpc.Dial(host)

	if err != nil {
		return nil, err
	}

	conn := ethclient.NewClient(ctx)
	return conn, nil
}

func parseBlockHeader(b *types.Block) (result GormModel.Block) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("[Error]", err)
			result = GormModel.Block{}
		}
	}()

	header := b.Header()
	result = GormModel.Block{
		Hash:             header.Hash().String(),
		ParentHash:       header.ParentHash.String(),
		Sha3Uncles:       header.UncleHash.String(),
		Miner:            header.Coinbase.String(),
		StateRoot:        header.Root.String(),
		TransactionsRoot: header.TxHash.String(),
		ReceiptsRoot:     header.ReceiptHash.String(),
		LogsBloom:        hex.EncodeToString(header.Bloom.Bytes()),
		Difficulty:       header.Difficulty.Int64(),
		Number:           header.Number.Int64(),
		GasLimit:         header.GasLimit,
		GasUsed:          header.GasUsed,
		Timestamp:        header.Time,
		MixHash:          header.MixDigest.String()}

	if header.BaseFee != nil {
		result.BaseFeePerGas = header.BaseFee.String()
	}

	extraDataBytes, err := hexutil.Bytes(header.Extra).MarshalText()
	if err == nil {
		result.ExtraData = string(extraDataBytes)
	}
	logsBloomBytes, err := hexutil.Bytes(header.Bloom[:]).MarshalText()
	if err == nil {
		result.LogsBloom = string(logsBloomBytes)
	}
	nonceByts, err := header.Nonce.MarshalText()
	if err == nil {
		result.Nonce = string(nonceByts)
	}
	return result
}

func parseTransaction(tx []*types.Transaction, blockNumber *int64) (result []GormModel.Transaction) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("[Error]", err)
			result = []GormModel.Transaction{}
		}
	}()

	for _, t := range tx {
		msg, err := t.AsMessage(types.LatestSignerForChainID(t.ChainId()), nil)
		if err != nil {
			return nil
		}
		v, r, s := t.RawSignatureValues()
		myTx := GormModel.Transaction{
			BlockNumber: *blockNumber,
			Hash:        t.Hash().String(),
			From:        msg.From().String(),
			To:          t.To().String(),
			Nonce:       t.Nonce(),
			GasPrice:    t.GasPrice().Int64(),
			Gas:         int64(t.Gas()),
			Value:       t.Value().Int64(),
			V:           v.Int64(),
			R:           r.Int64(),
			S:           s.Int64(),
		}
		inputBytes, err := hexutil.Bytes(t.Data()).MarshalText()
		if err == nil {
			myTx.Input = string(inputBytes)
		}
		result = append(result, myTx)
	}
	return
}

func getTransactionsReceiptLogs(client *ethclient.Client, txHash []common.Hash, blockNumber int64) (result []GormModel.ReceiptLog) {
	for _, h := range txHash {
		receipt, err := client.TransactionReceipt(context.Background(), h)
		if err != nil {
			LogHelper.LogFatal(fmt.Sprintf("getTransactionsReceipt, %v, %v", h, err))
			continue
		}

		for _, vLog := range receipt.Logs {

			l := GormModel.ReceiptLog{
				BlockNumber:     blockNumber,
				TransactionHash: h.String(),
				Index:           vLog.Index,
			}
			dataBytes, err := hexutil.Bytes(vLog.Data).MarshalText()
			if err == nil {
				l.Data = string(dataBytes)
			}
			result = append(result, l)
		}
	}
	return
}

func Login(r BaseModel.LoginRequest) (response BaseModel.LoginResponse) {
	if r.Account == "" || r.Password == "" {
		response.SetError(ResultCode.Parameter, "Missing Parameter.")
		return
	}

	//從資料庫取出密碼
	hashPwd := Repository.GetAccountPassword(r.Account)

	if hashPwd == nil {
		response.SetError(ResultCode.Parameter, "Login Failed, Account or Password Invalid.")
		return
	}

	//如果密碼不正確
	if !strings.EqualFold(*hashPwd, doMd5(r.Password)) {
		response.SetError(ResultCode.AccountNameIsAlreadyUsed, "Login Failed, Account or Password Invalid.")
		return
	}
	response.Data.AccessToken = createJwtToken(r.Account, []int{})

	return
}

func doMd5(input string) (output string) {
	data := []byte(input)
	has := md5.Sum(data)
	output = fmt.Sprintf("%x", has) //将[]byte转成16进制
	return
}

func createJwtToken(accountName string, functionIds []int) (output string) {
	jwtKey := []byte(ConfigHelper.GetString("JwtSettings.SignKey"))

	payload := BaseModel.JwtPayload{
		Account:     accountName,
		FunctionIds: functionIds,
		//+8時區
		Expires: time.Now().Add(time.Hour * time.Duration(ConfigHelper.GetInt64("JwtSettings.ExpireHour"))).UnixMilli(), //7天後jwt過期
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	output, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return
}

func VerifyJwtToken(jwtToken string, route string) (bool, string) {
	if jwtToken == "" {
		return false, "Token is empty."
	}
	var payload BaseModel.JwtPayload
	jwtKey := []byte(ConfigHelper.GetString("JwtSettings.SignKey"))

	token, err := jwt.ParseWithClaims(jwtToken, &payload, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			panic(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		return jwtKey, nil
	})

	if err != nil {
		panic(err)
	}

	if !token.Valid {
		return false, "Token Invalid."
	}

	//檢查時間是否過期
	if payload.Expires < time.Now().UnixMilli() {
		return false, "Token Invalid."
	}

	//檢查Token是否已登出
	// if Repository.IsJwtTokenLogout(jwtToken) {
	// 	return false, "Token Invalid."
	// }

	// //檢查權限足夠
	// routeFunctionId := Repository.GetRouteFunctionId(route)
	// if routeFunctionId == nil {
	// 	return false, "Route Premission Not Setting."
	// }

	// havePremission := false
	// for _, a := range payload.FunctionIds {
	// 	if a == *routeFunctionId {
	// 		havePremission = true
	// 		break
	// 	}
	// }

	// if !havePremission {
	// 	return false, "No Premission."
	// }

	return true, ""
}
