package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/y-ono366/percent-of-the-year-go/src/common"
	"github.com/y-ono366/percent-of-the-year-go/src/message"
)

type Response struct {
	Message string `json:"message"`
}

func Handler() (Response, error) {
	log := common.LogInit()
	parcent, isFlg := message.GetParcent()
	if isFlg == false {
		log.Println("1%の進捗なし")
		return Response{
			Message: "1%の進捗なし",
		}, nil
	}
	msg := message.CreateTweetMessage(parcent)
	fmt.Println(msg)
	return Response{
		Message: "完了",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
