package utils

import (
	"github.com/spf13/viper"
	"log"
)

// GetEnv .
func GetEnv(key string) string {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
