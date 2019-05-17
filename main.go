package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/y-ono366/percent-of-the-year-go/src/common"
	"github.com/y-ono366/percent-of-the-year-go/src/message"
	"github.com/y-ono366/percent-of-the-year-go/src/twitter"
	"log"
)

var Env string
var Log *log.Logger
var tw twitter.Twitter

func init() {
	Env = common.InitEnvironment()
	Log = common.LogInit()
}

func Handler() {
	parcent, isFlg := message.GetParcent()
	if isFlg == false {
		return
	}
	msg := message.CreateTweetMessage(parcent)
	tw.GetTwClient()
	tw.Post(msg)
}

func main() {
	if Env == "production" {
		lambda.Start(Handler)
	} else {
		if err := godotenv.Load(); err != nil {
			Log.Fatal(err)
		}
		Handler()
	}
}
