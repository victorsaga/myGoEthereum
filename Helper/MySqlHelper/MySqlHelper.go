package MySqlHelper

import (
	"fmt"
	"myGoEthereum/Helper/ConfigHelper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const CommonOffsetSql = " LIMIT ? OFFSET ? "
const NoRowsErrorMessage = "sql: no rows in result set"

func GetSqlxInstance() (db *sqlx.DB) {
	db = sqlx.MustConnect("mysql", getConnectString())
	return
}

func getConnectString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		ConfigHelper.GetString("DatabaseConnectInfo.UserName"),
		ConfigHelper.GetString("DatabaseConnectInfo.Password"),
		ConfigHelper.GetString("DatabaseConnectInfo.Url"),
		ConfigHelper.GetString("DatabaseConnectInfo.Port"),
		ConfigHelper.GetString("DatabaseConnectInfo.Database"))
}

func Convert_PageAndPageSize_ToMySqlLimitAndOffset(inputPage *int, inputPageSize *int) (limit int, offset int, outPage int, outPageSize int) {
	if inputPage == nil || *inputPage <= 0 {
		outPage = 1
	} else {
		outPage = *inputPage
	}

	if inputPageSize == nil || *inputPageSize <= 0 {
		outPageSize = 10
	} else {
		outPageSize = *inputPageSize
	}

	limit = outPageSize
	offset = outPageSize * (outPage - 1) //mysql是0為起點，page是1為起點，所以要-1
	return
}
