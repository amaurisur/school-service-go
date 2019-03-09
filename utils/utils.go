package utils

import (
	"os"
	"regexp"

	"github.com/astaxie/beego"
)

var configRex = regexp.MustCompile(`\{\{([\w_-]+)\}\}`)

// Getenv wraps getenv with a default value for the
// case when the variable is empty.
func Getenv(key, defaultValue string) string {
	if result := os.Getenv(key); len(result) != 0 {
		return result
	}

	return defaultValue
}

// GetAppName returns the service name from the config
func GetAppName() string {
	return GetAppConfig("appname", "")
}

// GetAppConfig :Get key configuration variable if exist otherwise return defaultValue
func GetAppConfig(key, defaultValue string) string {
	return beego.AppConfig.DefaultString(key, defaultValue)
}

// GetAppIntConfig :Get key configuration variable as Int if exist otherwise return defaultValue
// it is parsed from the secret manager if it's needed
func GetAppIntConfig(key string, defaultValue int) int {
	return beego.AppConfig.DefaultInt(key, defaultValue)
}
