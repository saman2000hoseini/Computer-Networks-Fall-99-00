package response

import (
	"encoding/json"
	"errors"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
)

const (
	SignType = "sign"
	Success  = "Welcome to our chat room"
)

type Sign struct {
	Message string `json:"message"`
}

func NewSignResponse(err error) *Sign {
	if err != nil {
		return &Sign{err.Error()}
	}

	return &Sign{Success}
}

func (s Sign) GenerateResponse() (*request.Request, error) {
	body, err := json.Marshal(s)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(SignType, body), nil
}
