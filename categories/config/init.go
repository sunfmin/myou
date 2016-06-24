package config

import (
	"log"
	"os"
	"strings"
)

// this is configured from env variables
var (
	Env               string
	MySQLHost         string
	MySQLPort         string
	MySQLRootPassword string
	MySQLDatabase     string
)

func init() {
	Env = envOrPanic("CATEGORIES_ENV", false)

	MySQLHost = envOrPanic("CATEGORIES_MYSQL_PORT_3306_TCP_ADDR", false)
	MySQLPort = envOrPanic("CATEGORIES_MYSQL_PORT_3306_TCP_PORT", false)
	MySQLRootPassword = envOrPanic("CATEGORIES_MYSQL_ENV_MYSQL_ROOT_PASSWORD", true)
	MySQLDatabase = envOrPanic("CATEGORIES_MYSQL_DATABASE", false)
}

func envOrPanic(key string, allowEmpty bool) (r string) {
	r = os.Getenv(key)
	if r == "" && !allowEmpty {
		panic("env " + key + " is not set")
	}
	logValue := r
	if strings.Contains(logValue, "PASSWORD") || strings.Contains(logValue, "SECRET") {
		logValue = "<HIDDEN>"
	}
	log.Printf("Configure: %s = %s\n", key, logValue)
	return
}
