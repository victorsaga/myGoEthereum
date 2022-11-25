package MySqlGormHelper

import (
	"fmt"
	"myGoEthereum/Helper/ConfigHelper"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DescString = "desc"
)

func GetGormInstance() (db *gorm.DB) {

	db, err := gorm.Open(mysql.Open(getConnectString()), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprint("connection to mysql failed:", err))
	}
	return
}

func IsNotFound(err error) bool {
	if err != nil && strings.EqualFold(err.Error(), "record not found") {
		return true
	}
	return false
}

func getConnectString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		ConfigHelper.GetString("DatabaseConnectInfo.UserName"),
		ConfigHelper.GetString("DatabaseConnectInfo.Password"),
		ConfigHelper.GetString("DatabaseConnectInfo.Url"),
		ConfigHelper.GetString("DatabaseConnectInfo.Port"),
		ConfigHelper.GetString("DatabaseConnectInfo.Database"))
}
