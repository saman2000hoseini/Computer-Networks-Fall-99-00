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
	Message     string   `json:"message"`
	OnlineUsers []string `json:"online_users"`
}

func NewSignResponse(err error, onlineUsers []string) *Sign {
	if err != nil {
		return &Sign{err.Error(), onlineUsers}
	}

	return &Sign{Success, onlineUsers}
}

func (s Sign) GenerateResponse() (*request.Request, error) {
	body, err := json.Marshal(s)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(SignType, body), nil
}

func LogOut(users *[]string, user string) {
	for index := range *users {
		if (*users)[index] == user {
			(*users)[index] = (*users)[len(*users)-1]
			(*users)[len(*users)-1] = ""
			*users = (*users)[:len(*users)-1]
			return
		}
	}
}
