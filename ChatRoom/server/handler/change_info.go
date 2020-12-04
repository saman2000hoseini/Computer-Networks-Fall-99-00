package handler

import (
	"encoding/json"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"time"
)

func (c *ClientHandler) HandleChangeInfo(body []byte, client *model.Client) error {
	info := &request.ChangeInfo{}
	err := json.Unmarshal(body, info)
	if err != nil {
		logrus.Errorf("change info handler: error while unmarshalling request: %s", err.Error())
		return err
	}

	stored, err := c.userRepo.Find(client.Username)
	if err != nil {
		req, err := response.NewSignResponse(err, nil).GenerateResponse()
		client.Out <- req

		return err
	}

	lastUser := stored.Username
	if len(info.Username) > 0 {
		stored.Username = info.Username
	}
	if len(info.Password) > 0 {
		stored.Password = info.Password
	}
	if len(info.Email) > 0 {
		stored.Email = info.Email
	}

	err = c.userRepo.Update(stored)
	if err != nil {
		req, _ := response.NewMessageResponse("server", client.Username,
			"error changing info: "+err.Error()).GenerateResponse()
		client.Out <- req
	} else if stored.Username != lastUser {
		client.Username = stored.Username
		c.clientsUser[stored.ID] = stored.Username
		delete(c.clientsID, lastUser)
		c.clientsID[stored.Username] = stored.ID
		c.informJoin(lastUser, false)
		req, _ := response.NewGlobalMessageResponse(newGlobalMsg(lastUser,
			"Changed username to => "+stored.Username), nil).GenerateResponse()
		c.informAll(req)

		req, _ = response.NewChangeInfoResponse(stored.Username).GenerateResponse()
		client.Out <- req

		<-time.Tick(10)
		c.informJoin(stored.Username, true)
	}

	return err
}
