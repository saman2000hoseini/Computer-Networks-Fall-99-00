package response

import (
	"encoding/json"
	"errors"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
)

const PrivateMessageType = "message"

type PrivateMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

func NewMessageResponse(from, to, message string) *PrivateMessage {
	return &PrivateMessage{
		From:    from,
		To:      to,
		Message: message,
	}
}

func (p PrivateMessage) GenerateResponse() (*request.Request, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(PrivateMessageType, body), nil
}
