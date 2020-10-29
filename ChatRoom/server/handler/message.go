package handler

import (
	"encoding/json"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/sirupsen/logrus"
)

func (c ClientHandler) HandlePrivateMessage(body []byte) {
	msg := &request.PrivateMessage{}
	err := json.Unmarshal(body, msg)
	if err != nil {
		logrus.Errorf("private message handler: err while unmarshalling request: %s", err.Error())
	}

}
