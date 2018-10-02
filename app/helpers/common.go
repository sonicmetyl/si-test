package helpers

import (
	"os"
	config "github.com/spf13/viper"
	"strconv"
)
func GetConfigString(key string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	} else if config.GetString(key) != "" {
		return config.GetString(key)
	}
	return ""
}

func GetConfigBool(key string) (exist bool) {
	if os.Getenv(key) != "" {
		b, err := strconv.ParseBool(key)
		if err != nil {
			return exist
		}
		return b
	} else if config.GetBool(key) {
		return config.GetBool(key)
	}
	return exist
}