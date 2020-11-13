package request

import (
	"encoding/json"
)

type Request struct {
	Type string `json:"type"`
	Body []byte `json:"body"`
}

func New(t string, rb []byte) *Request {
	return &Request{
		Type: t,
		Body: rb,
	}
}

func (r *Request) GenerateRequest() ([]byte, error) {
	return json.Marshal(r)
}
