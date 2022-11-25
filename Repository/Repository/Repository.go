package Repository

import (
	"fmt"
	"myGoEthereum/Helper/ConfigHelper"
	"myGoEthereum/Helper/MySqlGormHelper"
	"myGoEthereum/Model/GormModel"

	"gorm.io/gorm/clause"
)

func GetBlocksByLimit(limit int) (result []GormModel.Block) {
	db := MySqlGormHelper.GetGormInstance()
	db.Order(fmt.Sprintf("%s %s", GormModel.ColumnNameBlocksNumber, MySqlGormHelper.DescString)).Limit(limit).Find(&result)
	return
}

func GetBlocksByNumber(number int64) (result GormModel.Block) {
	db := MySqlGormHelper.GetGormInstance()
	db.Find(&result, number)
	return
}

func GetBlockTransactionHashes(blockNumber int64) (result []string) {
	db := MySqlGormHelper.GetGormInstance()
	db.Model(&GormModel.Transaction{}).Where(fmt.Sprintf("%s=?", GormModel.ColumnNameTransactionsBlockNumber), blockNumber).
		Pluck(GormModel.ColumnNameTransactionsHash, &result)
	return
}

func GetTransactionByHash(transactionHash string) (result GormModel.Transaction) {
	db := MySqlGormHelper.GetGormInstance()
	db.Find(&result, GormModel.ColumnNameTransactionsHash+"=?", transactionHash).Take(&result)
	return
}

func GetTransactionsReceiptLogs(transactionHash string) (result []GormModel.ReceiptLog) {
	db := MySqlGormHelper.GetGormInstance()
	db.Order(fmt.Sprintf("`%s`", GormModel.ColumnNameReceiptLogsIndex)).Find(&result, GormModel.ColumnNameReceiptLogsTransactionsHash+"=?", transactionHash)
	return
}

func GetDbMaxBlockNumber() int64 {
	var result []int64
	db := MySqlGormHelper.GetGormInstance()
	db.Model(&GormModel.Block{}).
		Select(fmt.Sprintf("ifnull(max(%s),0) as %s", GormModel.ColumnNameBlocksNumber, GormModel.ColumnNameBlocksNumber)).
		Pluck(GormModel.ColumnNameBlocksNumber, &result)
	return result[0]
}

func TruncateBlocksTransactionsReceiptLogs() {
	db := MySqlGormHelper.GetGormInstance()
	db.Exec("TRUNCATE TABLE  " + GormModel.TableNameBlocks)
	db.Exec("TRUNCATE TABLE  " + GormModel.TableNameTransactions)
	db.Exec("TRUNCATE TABLE  " + GormModel.TableNameReceiptLogs)
}

func InsertOnConflictUpdate(v interface{}) (rowsAffected int64, err error) {
	db := MySqlGormHelper.GetGormInstance()
	tx := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(v, ConfigHelper.GetInt("MySqlBatchSize"))

	if tx.RowsAffected > 0 {
		rowsAffected = tx.RowsAffected
	} else if tx.Error != nil && !MySqlGormHelper.IsNotFound(tx.Error) {
		panic(tx.Error)
	}
	return
}

func GetAccountPassword(accountName string) (hashPassword *string) {

	db := MySqlGormHelper.GetGormInstance()
	var account GormModel.Account
	tx := db.Take(&account, GormModel.ColumnNameAccountsName+"=?", accountName)

	if tx.RowsAffected == 1 {
		hashPassword = &account.Password
	} else if tx.Error != nil && !MySqlGormHelper.IsNotFound(tx.Error) {
		panic(tx.Error)
	}
	return
}

//sqlx
// func GetAccountPassword(accountName string) (hashPassword *string) {
// 	db := MySqlHelper.GetSqlxInstance()
// 	defer db.Close()

// 	err := db.Get(&hashPassword, `
// 		SELECT password
// 		FROM accounts
// 		WHERE enable = 1 AND name = ?`, accountName)

// 	if err != nil && err.Error() != MySqlHelper.NoRowsErrorMessage {
// 		panic(err)
// 	}

// 	return
// }
