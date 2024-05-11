package controller

import (
	"context"
	"encoding/json"
	"gocroot/config"
	"gocroot/helper"
	"gocroot/model"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var resp model.Response
	var err error
	if GetSecretFromHeader(request.Headers) == config.WebhookSecret {
		var msg model.IteungMessage
		err = json.Unmarshal([]byte(request.Body), &msg)
		if err != nil {
			resp.Response = err.Error()
		}
		resp = helper.WebHook(config.WAKeyword, config.WAPhoneNumber, config.WAAPIQRLogin, config.WAAPIMessage, msg, config.Mongoconn)
	} else {
		dt := &model.WebHook{
			URL:    config.WebhookURL,
			Secret: config.WebhookSecret,
		}
		res, err := helper.RefreshToken(dt, config.WAPhoneNumber, config.WAAPIGetToken, config.Mongoconn)
		if err != nil {
			resp.Response = err.Error()
		}
		resp.Response = jsonstr(res.ModifiedCount)
	}

	return events.APIGatewayProxyResponse{Body: jsonstr(resp), StatusCode: 200}, err
}

func jsonstr(strc interface{}) string {
	jsonData, err := json.Marshal(strc)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonData)
}

func GetSecretFromHeader(headers map[string]string) (secret string) {
	if headers["secret"] != "" {
		secret = headers["secret"]
	} else if headers["Secret"] != "" {
		secret = headers["Secret"]
	}
	return
}
