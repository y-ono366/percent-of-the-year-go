package common

import (
	"os"
)

var Env string

func InitEnvironment() string {
	switch os.Getenv("APP_ENV") {
	case "production":
		Env = "production"
	default:
		Env = "development"
	}
	return Env
}
func GetEnv() string {
	return Env
}
