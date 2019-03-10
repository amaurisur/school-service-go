package utils

import (
	"os"
	"regexp"
	"strconv"

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
	val := beego.AppConfig.DefaultString(key, defaultValue)
	match := configRex.FindStringSubmatch(val)
	if len(match) == 2 {
		return getEnv(match[1], defaultValue)
	}
	return val
}

// GetAppIntConfig :Get key configuration variable as Int if exist otherwise return defaultValue
// it is parsed from the secret manager if it's needed
func GetAppIntConfig(key string, defaultValue int) int {
	return beego.AppConfig.DefaultInt(key, defaultValue)
}

// getEnv get key environment variable if exist otherwise return defaultValue
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func GetAppBoolConfig(key string, defaultValue bool) bool {
	if b, err := strconv.ParseBool(GetAppConfig(key, strconv.FormatBool(defaultValue))); err != nil {
		return defaultValue
	} else {
		return b
	}
}
