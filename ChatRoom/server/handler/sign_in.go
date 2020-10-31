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

func (c *ClientHandler) HandleSignIn(body []byte, client *model.Client) (*request.Request, error) {
	info := &serverRequest.SignIn{}
	err := json.Unmarshal(body, info)
	if err != nil {
		logrus.Errorf("sign in handler: err while unmarshalling request: %s", err.Error())
		return nil, err
	}

	var stored model.User
	err = c.db.Where(&model.User{Username: info.Username}).First(&stored).Error
	if err != nil {
		return response.NewSignResponse(err, nil).GenerateResponse()
	}

	if !stored.CheckPassword(info.Password) {
		return response.NewSignResponse(errors.New(utils.ErrWrongPassword), nil).GenerateResponse()
	}

	client.Username = info.Username
	c.clients[info.Username] = client
	c.clientIDs = append(c.clientIDs, info.Username)
	c.informJoin(info.Username, true)
	return response.NewSignResponse(nil, c.clientIDs).GenerateResponse()
}
