package request

import (
	"encoding/json"
	"errors"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	"golang.org/x/crypto/bcrypt"
)

const SignInType = "signin"

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewSignInRequest(username, password string, email *string) (*SignIn, error) {
	if len(username) < 6 {
		return nil, errors.New("name is not under 6 chars")
	}
	if len(password) < 8 {
		return nil, errors.New("password is not under 8 chars")
	}

	hPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &SignIn{
		Username: username,
		Password: string(hPass),
	}, nil
}

func (s SignIn) GenerateRequest() (*request.Request, error) {
	body, err := json.Marshal(s)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(SignInType, body), nil
}
