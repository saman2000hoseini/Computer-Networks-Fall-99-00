package response

import (
	"encoding/json"
	"errors"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
)

const (
	SignType = "sign"
	Success  = "Welcome to our chat room"
)

type Sign struct {
	Message     string   `json:"message"`
	OnlineUsers []string `json:"online_users"`
}

func NewSignResponse(err error, onlineUsers map[string]uint64) *Sign {
	users := make([]string, 0, len(onlineUsers))
	for k := range onlineUsers {
		users = append(users, k)
	}

	if err != nil {
		return &Sign{err.Error(), users}
	}

	return &Sign{Success, users}
}

func (s Sign) GenerateResponse() (*request.Request, error) {
	body, err := json.Marshal(s)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(SignType, body), nil
}

func LogOut(users *[]uint64, id uint64) {
	for index := range *users {
		if (*users)[index] == id {
			(*users)[index] = (*users)[len(*users)-1]
			(*users)[len(*users)-1] = 0
			*users = (*users)[:len(*users)-1]
			return
		}
	}
}
