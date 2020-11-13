package request

import (
	"encoding/json"
	"errors"
)

const SignInType = "signin"

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewSignInRequest(username, password string) (*SignIn, error) {
	if len(username) < 6 {
		return nil, errors.New("name is not under 6 chars")
	}
	if len(password) < 8 {
		return nil, errors.New("password is not under 8 chars")
	}

	return &SignIn{
		Username: username,
		Password: password,
	}, nil
}

func (s SignIn) GenerateRequest() (*Request, error) {
	body, err := json.Marshal(s)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(SignInType, body), nil
}
