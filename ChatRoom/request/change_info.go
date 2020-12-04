package request

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

const ChangeInfoType = "change_info"

type ChangeInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewChangeInfoRequest(username, password string, email string) (*ChangeInfo, error) {
	if len(username) < 6 && len(password) < 8 && len(email) < 6 {
		return nil, errors.New("please fill at least one field")
	}

	if len(username) > 0 && len(username) < 6 {
		return nil, errors.New("name can't be under 6 chars")
	}
	if len(password) > 0 && len(password) < 8 {
		return nil, errors.New("password can't be under 8 chars")
	}

	if len(password) > 0 {
		hPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		password = string(hPass)
	}

	return &ChangeInfo{
		Username: username,
		Password: password,
		Email:    email,
	}, nil
}

func (c ChangeInfo) GenerateRequest() (*Request, error) {
	body, err := json.Marshal(c)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(ChangeInfoType, body), nil
}
