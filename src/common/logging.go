package common

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func LogInit() *logrus.Logger {
	log := logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	Log = log
	return Log

}

func GetLog() *logrus.Logger {
	return Log
}
