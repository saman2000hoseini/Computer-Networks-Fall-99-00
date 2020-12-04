package handler

import (
	"encoding/json"
	"errors"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
)

func (c *ClientHandler) HandleSignIn(body []byte, client *model.Client) error {
	info := &request.SignIn{}
	err := json.Unmarshal(body, info)
	if err != nil {
		logrus.Errorf("sign in handler: err while unmarshalling request: %s", err.Error())
		return err
	}

	stored, err := c.userRepo.Find(info.Username)
	if err != nil {
		req, err := response.NewSignResponse(err, nil).GenerateResponse()
		client.Out <- req

		return err
	}

	if !stored.CheckPassword(info.Password) {
		req, err := response.NewSignResponse(errors.New(pkg.ErrWrongPassword), nil).GenerateResponse()
		client.Out <- req

		return err
	}

	client.Username = stored.Username
	client.ID = stored.ID
	c.clients[stored.ID] = client
	c.clientIDs = append(c.clientIDs, stored.ID)
	c.clientsUser[stored.ID] = stored.Username
	c.clientsID[stored.Username] = stored.ID
	c.informJoin(stored.Username, true)

	req, err := response.NewSignResponse(nil, c.clientsID).GenerateResponse()
	client.Out <- req

	return err
}
