package handler

import (
	"encoding/json"
	"errors"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/utils"
	serverRequest "github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
)

func (c ClientHandler) HandleSignIn(body []byte, client *model.Client) (*request.Request, error) {
	info := &serverRequest.SignIn{}
	err := json.Unmarshal(body, info)
	if err != nil {
		logrus.Errorf("sign in handler: err while unmarshalling request: %s", err.Error())
		return nil, err
	}

	var stored model.User
	err = c.db.Where(&model.User{Username: info.Username}).First(&stored).Error
	if err != nil {
		return response.NewSignResponse(err).GenerateResponse()
	}

	if info.Password != stored.Password {
		return response.NewSignResponse(errors.New(utils.ErrWrongPassword)).GenerateResponse()
	}

	c.clients[info.Username] = client
	return response.NewSignResponse(nil).GenerateResponse()
}
