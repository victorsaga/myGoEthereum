package MySqlGormHelper

import (
	"fmt"
	"myGoEthereum/Helper/ConfigHelper"
	"strings"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DescString = "desc"
)

var gormMutex sync.Mutex
var db *gorm.DB

func GetGormInstance() *gorm.DB {
	if db == nil {
		gormMutex.Lock()
		defer gormMutex.Unlock()
		var err error
		if db == nil {
			db, err = gorm.Open(mysql.Open(getConnectString()), &gorm.Config{})

			if err != nil {
				panic(fmt.Sprint("connection to mysql failed:", err))
			}
		}
	}
	return db
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
