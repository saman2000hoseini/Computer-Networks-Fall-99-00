package handler

import (
	"encoding/json"
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/utils"
	serverRequest "github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
)

func (c *ClientHandler) HandlePrivateMessage(body []byte, client *model.Client) error {
	msg := &serverRequest.PrivateMessage{}
	err := json.Unmarshal(body, msg)
	if err != nil {
		logrus.Errorf("private message handler: err while unmarshalling request: %s", err.Error())
		return err
	}

	if msg.To == "all" {
		req, err := response.NewGlobalMessageResponse(newGlobalMsg(msg.From,
			msg.Message), nil).GenerateResponse()

		c.informAll(req)
		return err
	}

	fmt.Printf("%s, %v\n", msg.To, c.clients[msg.To])
	if c.clients[msg.To] != nil {
		req, err := response.NewMessageResponse(msg.From, msg.To, msg.Message).GenerateResponse()
		c.clients[msg.To].Out <- req
		return err
	}

	msg.To = ""
	msg.Message = utils.ErrUserNotFound
	req, err := response.NewMessageResponse(msg.From, msg.To, msg.Message).GenerateResponse()

	client.Out <- req

	return err
}

func newGlobalMsg(from, msg string) string {
	return fmt.Sprintf("Global: %s: %s", from, msg)
}
