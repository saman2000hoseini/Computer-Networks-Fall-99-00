package request

import (
	"encoding/json"
	"errors"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	"strings"
)

const PrivateMessageType = "message"

type PrivateMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

func NewMessageRequest(from, to, message string) (*PrivateMessage, error) {
	message = strings.TrimSpace(message)
	if len(message) < 1 {
		return nil, errors.New("message cannot be empty")
	}

	return &PrivateMessage{
		From:    from,
		To:      to,
		Message: message,
	}, nil
}

func (p PrivateMessage) GenerateRequest() (*request.Request, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(PrivateMessageType, body), nil
}
