package response

import (
	"encoding/json"
	"errors"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
)

const GlobalMessageType = "info"

type GlobalMessage struct {
	Message string `json:"message"`
	Joined  *bool
}

func NewGlobalMessageResponse(message string, joined *bool) *GlobalMessage {
	return &GlobalMessage{
		Message: message,
		Joined:  joined,
	}
}

func (p GlobalMessage) GenerateResponse() (*request.Request, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(GlobalMessageType, body), nil
}
