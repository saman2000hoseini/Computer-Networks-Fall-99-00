package response

import (
	"encoding/json"
	"errors"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
)

const ChangeUserType = "change_user"

type ChangeUser struct {
	Username string `json:"username"`
}

func NewChangeInfoResponse(username string) *ChangeUser {
	return &ChangeUser{
		Username: username,
	}
}

func (c ChangeUser) GenerateResponse() (*request.Request, error) {
	body, err := json.Marshal(c)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(ChangeUserType, body), nil
}
