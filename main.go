package main

import (
	"context"
	"fmt"
	"gocroot/model"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event *model.IteungMessage) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	message := fmt.Sprintf("Hello %s!", event.Alias_name)
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
