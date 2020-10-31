package handler

import (
	"encoding/json"
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/sirupsen/logrus"
)

func (c *ClientHandler) HandleSign(body []byte) error {
	resp := &response.Sign{}
	err := json.Unmarshal(body, resp)
	if err != nil {
		logrus.Errorf("client sign in handler: err while unmarshalling response: %s", err.Error())
		c.waiter <- false
		c.username = nil
		return err
	}

	if resp.Message == response.Success {
		c.signedIn = true
		c.Users = resp.OnlineUsers
		c.waiter <- true
		c.usersChange <- true
		return nil
	}

	fmt.Println(resp.Message)
	c.waiter <- true

	return nil
}
