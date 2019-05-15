package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/y-ono366/percent-of-the-year-go/common"
)

type Response struct {
	Message string `json:"message"`
}

func Handler() (Response, error) {
	log := common.LogInit()
	fmt.Println("fmt not found request_token_secret")
	log.Fatal("not found request_token_secret")
	return Response{
		Message: "Go Serverless v1.0! Your function executed successfully!",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
