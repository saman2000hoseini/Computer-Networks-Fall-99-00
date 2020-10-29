package request

import "github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"

type Request interface {
	GenerateRequest() (*request.Request, error)
}

func New(t string, rb []byte) *request.Request {
	return &request.Request{
		Type: t,
		Body: rb,
	}
}
