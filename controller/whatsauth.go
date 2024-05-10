package controller

import (
	"gocroot/config"
	"gocroot/helper"
	"gocroot/model"
	"gocroot/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func WhatsAuthReceiver(msg model.IteungMessage) (resp model.Response) {
	if pkg.IsLoginRequest(msg, config.WAKeyword) { //untuk whatsauth request login
		resp = pkg.HandlerQRLogin(msg, config.WAKeyword)
	} else { //untuk membalas pesan masuk
		resp = pkg.HandlerIncomingMessage(msg)
	}
	return
}

func RefreshWAToken() (res *mongo.UpdateResult, err error) {
	dt := &model.WebHook{
		URL:    config.WebhookURL,
		Secret: config.WebhookSecret,
	}
	resp, err := helper.PostStructWithToken[model.User]("Token", pkg.WAAPIToken(config.WAPhoneNumber), dt, config.WAAPIGetToken)
	if err != nil {
		return
	}
	profile := &model.Profile{
		Phonenumber: resp.PhoneNumber,
		Token:       resp.Token,
	}
	res, err = helper.ReplaceOneDoc(config.Mongoconn, "profile", bson.M{"phonenumber": resp.PhoneNumber}, profile)
	if err != nil {
		return
	}
	return
}
