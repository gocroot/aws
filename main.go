package main

import (
	"gocroot/controller"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(controller.HandleRequest)
}
