package handler

import (
	"encoding/json"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/utils"
	serverRequest "github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
)

func (c *ClientHandler) HandlePrivateMessage(body []byte, client *model.Client) (*model.Client, *request.Request, error) {
	msg := &serverRequest.PrivateMessage{}
	err := json.Unmarshal(body, msg)
	if err != nil {
		logrus.Errorf("private message handler: err while unmarshalling request: %s", err.Error())
		return client, nil, err
	}

	if c.clients[msg.To] != nil {
		client = c.clients[msg.To]
	} else {
		msg.To = ""
		msg.Message = utils.ErrUserNotFound
	}

	req, err := response.NewMessageResponse(msg.From, msg.To, msg.Message).GenerateResponse()
	return client, req, err
}
