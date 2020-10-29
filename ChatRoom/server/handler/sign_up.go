package handler

import (
	"encoding/json"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	serverRequest "github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
)

func (c ClientHandler) HandleSignUp(body []byte, client *model.Client) (*request.Request, error) {
	info := &serverRequest.SignUp{}
	err := json.Unmarshal(body, info)
	if err != nil {
		logrus.Errorf("sign up handler: err while unmarshalling request: %s", err.Error())
	}

	err = c.db.Create(info).Error
	if err == nil {
		c.clients[info.Username] = client
	}

	return response.NewSignResponse(err).GenerateResponse()
}
