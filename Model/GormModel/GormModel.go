package GormModel

import (
	"sync"
)

type Account struct {
	Name     string `gorm:"column:name;primaryKey"`
	Password string `gorm:"column:password;NOT NULL"`
	Enable   int8   `gorm:"column:enable;NOT NULL"`
}

type BlocksTransactionsMutex struct {
	Blocks       []Block
	Transactions []Transaction
	ReceiptLogs  []ReceiptLog
	Mux          sync.Mutex
}

const (
	TableNameBlocks       = "blocks"
	TableNameTransactions = "transactions"
	TableNameReceiptLogs  = "receipt_logs"

	ColumnNameBlocksNumber                = "number"
	ColumnNameTransactionsBlockNumber     = "blockNumber"
	ColumnNameTransactionsHash            = "hash"
	ColumnNameReceiptLogsTransactionsHash = "transactionHash"
	ColumnNameReceiptLogsIndex            = "index"
	ColumnNameAccountsName                = "name"
)

type Block struct {
	Number           int64  `gorm:"column:number;primary_key" json:"block_num"`
	Hash             string `gorm:"column:hash;unique;NOT NULL" json:"block_hash"`
	BaseFeePerGas    string `gorm:"column:baseFeePerGas" json:"-"`
	Difficulty       int64  `gorm:"column:difficulty;NOT NULL" json:"-"`
	ExtraData        string `gorm:"column:extraData;NOT NULL" json:"-"`
	GasLimit         uint64 `gorm:"column:gasLimit;NOT NULL" json:"-"`
	GasUsed          uint64 `gorm:"column:gasUsed;NOT NULL" json:"-"`
	LogsBloom        string `gorm:"column:logsBloom;NOT NULL" json:"-"`
	Miner            string `gorm:"column:miner;NOT NULL" json:"-"`
	MixHash          string `gorm:"column:mixHash;NOT NULL" json:"-"`
	Nonce            string `gorm:"column:nonce;NOT NULL" json:"-"`
	Timestamp        uint64 `gorm:"column:timestamp;NOT NULL" json:"block_time"`
	ParentHash       string `gorm:"column:parentHash;NULL" json:"parent_hash"`
	ReceiptsRoot     string `gorm:"column:receiptsRoot;NOT NULL" json:"-"`
	Sha3Uncles       string `gorm:"column:sha3Uncles;NULL" json:"-"`
	StateRoot        string `gorm:"column:stateRoot;NOT NULL" json:"-"`
	TransactionsRoot string `gorm:"column:transactionsRoot;NOT NULL" json:"-"`
}

type Transaction struct {
	BlockNumber      int64  `gorm:"column:blockNumber;primaryKey" json:"-"`
	Hash             string `gorm:"column:hash;primaryKey" json:"tx_hash"`
	From             string `gorm:"column:from;NOT NULL" json:"from"`
	To               string `gorm:"column:to;NOT NULL" json:"to"`
	Gas              int64  `gorm:"column:gas;NOT NULL" json:"-"`
	GasPrice         int64  `gorm:"column:gasPrice;NOT NULL" json:"-"`
	Nonce            uint64 `gorm:"column:nonce;NOT NULL" json:"nonce"`
	Input            string `gorm:"column:input;NOT NULL" json:"data"`
	TransactionIndex int64  `gorm:"column:transactionIndex;NOT NULL" json:"-"`
	Value            int64  `gorm:"column:value;NOT NULL" json:"value"`
	Type             int64  `gorm:"column:type;NOT NULL" json:"-"`
	V                int64  `gorm:"column:v;NOT NULL" json:"-"`
	R                int64  `gorm:"column:r;NOT NULL" json:"-"`
	S                int64  `gorm:"column:s;NOT NULL" json:"-"`
}

type ReceiptLog struct {
	BlockNumber     int64  `gorm:"column:blockNumber;primaryKey" json:"-"`
	TransactionHash string `gorm:"column:transactionHash;primaryKey" json:"-"`
	Index           uint   `gorm:"column:index;primaryKey" json:"index"`
	Data            string `gorm:"column:data;NOT NULL" json:"data"`
}
type BlocksResponse struct {
	Blocks []BlockResponseData `json:"blocks"`
}

type BlockResponseData struct {
	Block
	Stable bool `json:"Stable"`
}

type BlockTransactionsResponse struct {
	BlockResponseData
	Transactions []string `json:"transactions"`
}

type TransactionReceiptLogsResponse struct {
	Transaction
	Logs []ReceiptLog
}
