package handler

import (
	"encoding/json"
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/sirupsen/logrus"
)

func (c *ClientHandler) HandleGlobalMessage(body []byte) error {
	resp := &response.GlobalMessage{}
	err := json.Unmarshal(body, resp)
	if err != nil {
		logrus.Errorf("client global message handler: error while unmarshalling response: %s", err.Error())
		return err
	}

	if resp.Joined != nil {
		if *resp.Joined {
			c.Users = append(c.Users, resp.Message)
			c.messages <- fmt.Sprintf("\x1b[0;32m+ %s connected\033[0m\n", resp.Message)
		} else {
			response.LogOut(c.Users, resp.Message)
			c.messages <- fmt.Sprintf("\x1b[0;31m- %s disconnected\033[0m\n", resp.Message)
		}
		c.usersChange <- true
	}

	fmt.Println(resp.Message)

	c.waiter <- true

	return nil
}
