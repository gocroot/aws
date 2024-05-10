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
	fmt.Println(request.Headers["secret"])
	key1 := "secret"
	key2 := "Secret"
	if value, exists := request.Headers[key1]; exists {
		fmt.Printf("Key '%s' found with value '%s'\n", key1, value)
	} else if value, exists := request.Headers[key2]; exists {
		fmt.Printf("Key '%s' found with value '%s'\n", key1, value)
	}

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}
