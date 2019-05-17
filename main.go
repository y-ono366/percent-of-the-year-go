package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/y-ono366/percent-of-the-year-go/src/common"
	"github.com/y-ono366/percent-of-the-year-go/src/message"
	"github.com/y-ono366/percent-of-the-year-go/src/twitter"
)

var Env string
var Log *logrus.Logger
var tw twitter.Twitter

func init() {
	Env = common.InitEnvironment()
	Log = common.LogInit()
}

func Handler() {
	Log.Info("開始")
	parcent, isFlg := message.GetParcent()
	if isFlg == false {
		Log.Info("終了 進捗無し")
		return
	}
	msg := message.CreateTweetMessage(parcent)
	tw.GetTwClient()
	tw.Post(msg)
	Log.Info("終了 投稿完了")
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
