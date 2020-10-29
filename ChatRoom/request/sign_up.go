package request

import (
	"encoding/json"
	"errors"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	"golang.org/x/crypto/bcrypt"
)

const SignUpType = "signup"

type SignUp struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    *string `json:"email"`
}

func NewSignUpRequest(username, password string, email *string) (*SignUp, error) {
	if len(username) < 6 {
		return nil, errors.New("name can't be under 6 chars")
	}
	if len(password) < 8 {
		return nil, errors.New("password can't be under 8 chars")
	}

	hPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &SignUp{
		Username: username,
		Password: string(hPass),
		Email:    email,
	}, nil
}

func (s SignUp) GenerateRequest() (*request.Request, error) {
	body, err := json.Marshal(s)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(SignUpType, body), nil
}

type User interface {
	Create(username, password, email string)
	Update(username, password, newPassword, email *string)
	Delete(username, password string)
	Find(username string) uint
	AddToGroup(username, groupID string)
}
