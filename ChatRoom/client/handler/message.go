package handler

import (
	"encoding/json"
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/sirupsen/logrus"
)

func (c *ClientHandler) HandlePrivateMessage(body []byte) error {
	msg := &request.PrivateMessage{}
	err := json.Unmarshal(body, msg)
	if err != nil {
		logrus.Errorf("client private message handler: err while unmarshalling message body: %s", err.Error())
		return err
	}

	fmt.Printf("%s: %s\n", msg.From, msg.Message)

	return nil
}
