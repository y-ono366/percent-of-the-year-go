package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/y-ono366/percent-of-the-year-go/src/common"
)

type Response struct {
	Message string `json:"message"`
}

func Handler() (Response, error) {
	log := common.LogInit()
	log.Print("test")
	return Response{
		Message: "Go Serverless v1.0! Your function executed successfully!",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
