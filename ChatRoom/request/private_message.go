package request

import (
	"encoding/json"
	"errors"
	"strings"
)

const message = "message"

type PrivateMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

func NewMessageRequest(from, to, message string) (*PrivateMessage, error) {
	message = strings.TrimSpace(message)
	if len(message) < 1 {
		return nil, errors.New("name is not under 6 chars")
	}

	return &PrivateMessage{
		From:    from,
		To:      to,
		Message: message,
	}, nil
}

func (s PrivateMessage) GenerateRequest() (*PackedRequest, error) {
	body, err := json.Marshal(s)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(message, body), nil
}
