package common

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func LogInit() *logrus.Logger {
	log := logrus.New()
	Log = log
	return Log
}

func GetLog() *logrus.Logger {
	return Log
}
