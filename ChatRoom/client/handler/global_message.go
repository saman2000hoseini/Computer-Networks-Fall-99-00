package handler

import (
	"encoding/json"
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/sirupsen/logrus"
	"time"
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
			c.messages <- fmt.Sprintf("[%v] \033[1;32m+ %s connected\033[0m", time.Now().Local(), resp.Message)
		} else {
			logOut(&c.Users, resp.Message)
			c.messages <- fmt.Sprintf("[%v] \x1b[0;31m- %s disconnected\033[0m", time.Now().Local(), resp.Message)
		}
		c.usersChange <- true
	} else {
		c.messages <- fmt.Sprintf("[%v] \033[1;34m %s \033[0m", time.Now().Local(), resp.Message)
	}

	return nil
}

func logOut(users *[]string, id string) {
	for index := range *users {
		if (*users)[index] == id {
			(*users)[index] = (*users)[len(*users)-1]
			(*users)[len(*users)-1] = ""
			*users = (*users)[:len(*users)-1]
			return
		}
	}
}
