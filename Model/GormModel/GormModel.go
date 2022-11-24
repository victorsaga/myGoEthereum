package GormModel

import (
	"sync"
	"time"
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

	ColumnNameBlocksNumber = "number"
)

type Block struct {
	Number           int64     `gorm:"column:number;primary_key"`
	Hash             string    `gorm:"column:hash;unique;NOT NULL"`
	BaseFeePerGas    string    `gorm:"column:baseFeePerGas"`
	Difficulty       int64     `gorm:"column:difficulty;NOT NULL"`
	ExtraData        string    `gorm:"column:extraData;NOT NULL"`
	GasLimit         uint64    `gorm:"column:gasLimit;NOT NULL"`
	GasUsed          uint64    `gorm:"column:gasUsed;NOT NULL"`
	LogsBloom        string    `gorm:"column:logsBloom;NOT NULL"`
	Miner            string    `gorm:"column:miner;NOT NULL"`
	MixHash          string    `gorm:"column:mixHash;NOT NULL"`
	Nonce            string    `gorm:"column:nonce;NOT NULL"`
	ParentHash       string    `gorm:"column:parentHash;NULL"`
	ReceiptsRoot     string    `gorm:"column:receiptsRoot;NOT NULL"`
	Sha3Uncles       string    `gorm:"column:sha3Uncles;NULL"`
	StateRoot        string    `gorm:"column:stateRoot;NOT NULL"`
	Timestamp        uint64    `gorm:"column:timestamp;NOT NULL"`
	TransactionsRoot string    `gorm:"column:transactionsRoot;NOT NULL"`
	UpdateAt         time.Time `gorm:"column:updateAt;default:CURRENT_TIMESTAMP"`
}

type Transaction struct {
	BlockNumber      int64  `gorm:"column:blockNumber;primaryKey"`
	Hash             string `gorm:"column:hash;primaryKey"`
	From             string `gorm:"column:from;NOT NULL"`
	Gas              int64  `gorm:"column:gas;NOT NULL"`
	GasPrice         int64  `gorm:"column:gasPrice;NOT NULL"`
	Input            string `gorm:"column:input;NOT NULL"`
	Nonce            uint64 `gorm:"column:nonce;NOT NULL"`
	To               string `gorm:"column:to;NOT NULL"`
	TransactionIndex int64  `gorm:"column:transactionIndex;NOT NULL"`
	Value            int64  `gorm:"column:value;NOT NULL"`
	Type             int64  `gorm:"column:type;NOT NULL"`
	V                int64  `gorm:"column:v;NOT NULL"`
	R                int64  `gorm:"column:r;NOT NULL"`
	S                int64  `gorm:"column:s;NOT NULL"`
}

type ReceiptLog struct {
	BlockNumber     int64  `gorm:"column:blockNumber;primaryKey"`
	TransactionHash string `gorm:"column:transactionHash;primaryKey"`
	Index           uint   `gorm:"column:index;primaryKey"`
	Data            string `gorm:"column:data;NOT NULL"`
}
