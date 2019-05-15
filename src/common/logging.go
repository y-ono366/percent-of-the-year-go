package common

import (
	"log"
	"os"
)

var Log *log.Logger

func LogInit() *log.Logger {
	logging := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	logging.SetFlags(log.LstdFlags | log.Lshortfile)
	Log = logging
	return Log
}

func GetLog() *log.Logger {
	return Log
}
