package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	serverRequest "github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func (c *ClientHandler) HandlePrivateMessage(body []byte) error {
	msg := &serverRequest.PrivateMessage{}
	err := json.Unmarshal(body, msg)
	if err != nil {
		logrus.Errorf("client private message handler: err while unmarshalling message body: %s", err.Error())
		return err
	}

	c.messages <- msgToString(msg)
	return nil
}

func (c *ClientHandler) Send(g *gocui.Gui, v *gocui.View) error {
	msg, req, err := c.parseInput(*c.username, v.Buffer())
	if err != nil {
		c.messages <- err.Error()
		return err
	}

	if msg.To != "all" {
		c.messages <- msgToString(msg)
	}

	c.client.Out <- req
	g.Update(func(g *gocui.Gui) error {
		v.Clear()
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)
		return nil
	})
	return nil
}

func (c *ClientHandler) parseInput(username, input string) (*serverRequest.PrivateMessage, *request.Request, error) {
	args := strings.Split(input, ">")
	args[0] = strings.TrimSpace(args[0])
	switch args[0] {
	default:
		if args[0] != "all" && !c.contains(args[0]) {
			return nil, nil, errors.New("user does not exist")
		}
	}

	msg, _ := serverRequest.NewMessageRequest(username, args[0], args[1])

	req, err := msg.GenerateRequest()

	return msg, req, err
}

func msgToString(msg *serverRequest.PrivateMessage) string {
	return fmt.Sprintf("[%v] %s: %s", time.Now().Local(), msg.From, msg.Message)
}

func (c *ClientHandler) contains(username string) bool {
	for i := range c.Users {
		if c.Users[i] == username {
			return true
		}
	}

	return false
}
