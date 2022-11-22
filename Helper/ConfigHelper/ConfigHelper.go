package ConfigHelper

import (
	"github.com/spf13/viper"
)

const RequestGuidKey = "RequestGuid"

func init() {
	viper.SetConfigName("Config")
	viper.AddConfigPath("./")
	viper.ReadInConfig()
}

func GetString(str string) string {
	return viper.GetString(str)
}

func GetInt64(str string) int64 {
	return viper.GetInt64(str)
}

func GetInt(str string) int {
	return viper.GetInt(str)
}

func SetString(key string, value string) {
	viper.Set(key, value)
}
