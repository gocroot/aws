package controller

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}
	fmt.Println("apakah ada secret?")
	var secret string
	if request.Headers["secret"] != "" {
		secret = request.Headers["secret"]
	} else if request.Headers["Secret"] != "" {
		secret = request.Headers["secret"]
	}
	fmt.Println(secret)
	fmt.Println("diatas secretnya cuk")
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}
