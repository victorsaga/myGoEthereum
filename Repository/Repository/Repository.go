package Repository

import (
	"myGoEthereum/Helper/MySqlGormHelper"
	"myGoEthereum/Model/GormModel"
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

func GetAccountPassword(accountName string) (hashPassword *string) {

	db := MySqlGormHelper.GetGormInstance()
	var account GormModel.Account
	tx := db.Where("name=?", accountName).First(&account)

	if tx.RowsAffected == 1 {
		hashPassword = &account.Password
	} else if tx.Error != nil && !MySqlGormHelper.IsNotFound(tx.Error) {
		panic(tx.Error)
	}
	return
}
