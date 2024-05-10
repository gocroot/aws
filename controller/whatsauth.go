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
	fmt.Println(GetSecretFromHeader(request.Headers))
	fmt.Println("diatas secretnya cuk")
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func GetSecretFromHeader(headers map[string]string) (secret string) {
	if headers["secret"] != "" {
		secret = headers["secret"]
	} else if headers["Secret"] != "" {
		secret = headers["Secret"]
	}
	return
}
