package handler

import (
	"encoding/json"
	"fmt"
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

	if msg.To == "all" {
		req, err := response.NewGlobalMessageResponse(newGlobalMsg(msg.From,
			msg.Message), nil).GenerateResponse()

		c.informAll(req)
		return nil, nil, err
	}

	fmt.Printf("%s, %v\n", msg.To, c.clients[msg.To])
	if c.clients[msg.To] != nil {
		req, err := response.NewMessageResponse(msg.From, msg.To, msg.Message).GenerateResponse()
		return c.clients[msg.To], req, err
	}

	msg.To = ""
	msg.Message = utils.ErrUserNotFound
	req, err := response.NewMessageResponse(msg.From, msg.To, msg.Message).GenerateResponse()
	return client, req, err
}

func newGlobalMsg(from, msg string) string {
	return fmt.Sprintf("%s: %s", from, msg)
}
