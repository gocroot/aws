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
	key := "Secret"
	if value, exists := request.Headers[key]; exists {
		fmt.Printf("Key '%s' found with value '%s'\n", key, value)
	} else {
		fmt.Printf("Key '%s' not found\n", key)
	}

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}
