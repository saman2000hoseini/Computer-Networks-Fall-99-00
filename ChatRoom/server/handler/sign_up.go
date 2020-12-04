package handler

import (
	"encoding/json"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
)

func (c *ClientHandler) HandleSignUp(body []byte, client *model.Client) error {
	info := &request.SignUp{}
	err := json.Unmarshal(body, info)
	if err != nil {
		logrus.Errorf("sign up handler: err while unmarshalling request: %s", err.Error())
		return err
	}

	user := model.NewUser(info.Username, info.Password, info.Email)

	err = c.userRepo.Save(user)
	if err == nil {
		client.Username = info.Username
		c.clients[info.Username] = client
		c.clientIDs = append(c.clientIDs, info.Username)
		c.informJoin(info.Username, true)
	}

	req, err := response.NewSignResponse(err, c.clientIDs).GenerateResponse()
	client.Out <- req

	return err
}
