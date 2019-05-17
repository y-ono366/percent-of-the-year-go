package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/y-ono366/percent-of-the-year-go/src/common"
	"github.com/y-ono366/percent-of-the-year-go/src/message"
	"github.com/y-ono366/percent-of-the-year-go/src/twitter"
	"log"
	"os"
)

var Env string
var Log *log.Logger

type Response struct {
	Message string `json:"message"`
}

func init() {
	setEnvironment()
	Log = common.LogInit()
}

func setEnvironment() {
	switch os.Getenv("APP_ENV") {
	case "production":
		Env = "production"
	default:
		Env = "development"
	}
}

func Handler() (Response, error) {
	parcent, isFlg := message.GetParcent()
	if isFlg == false {
		Log.Println("1%の進捗なし")
		return Response{
			Message: "1%の進捗なし",
		}, nil
	}
	msg := message.CreateTweetMessage(parcent)
	fmt.Println(msg)
	client, err := twitter.GetTwClient()
	if err != nil {
	}
	fmt.Println(client)

	return Response{
		Message: "完了",
	}, nil
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
