package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Error read config")
	}
}

//GetString to get string from config.json
func GetString(key string) string {
	return viper.GetString(key)
}

//GetInt to get int from config.json
func GetInt(key string) int {
	return viper.GetInt(key)
}
