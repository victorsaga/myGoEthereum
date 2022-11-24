package Repository

import (
	"fmt"
	"myGoEthereum/Helper/ConfigHelper"
	"myGoEthereum/Helper/MySqlGormHelper"
	"myGoEthereum/Model/GormModel"

	"gorm.io/gorm/clause"
)

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

func GetDbMaxBlockNumber() int64 {
	var result []int64
	db := MySqlGormHelper.GetGormInstance()
	db.Table(GormModel.TableNameBlocks).
		Select(fmt.Sprintf("max(%s) as %s", GormModel.ColumnNameBlocksNumber, GormModel.ColumnNameBlocksNumber)).
		Pluck(GormModel.ColumnNameBlocksNumber, &result)
	if len(result) > 0 {
		return result[0]
	}
	return 0
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
	tx := db.Take(&account, "name=?", accountName)

	if tx.RowsAffected == 1 {
		hashPassword = &account.Password
	} else if tx.Error != nil && !MySqlGormHelper.IsNotFound(tx.Error) {
		panic(tx.Error)
	}
	return
}
