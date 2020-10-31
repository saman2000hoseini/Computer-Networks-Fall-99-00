package handler

import (
	"encoding/json"
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

	fmt.Sprintf("[%v] %s: %s", time.Now().Local(), msg.From, msg.Message)
	return nil
}

func (c *ClientHandler) Send(g *gocui.Gui, v *gocui.View) error {
	req, err := parseInput(*c.username, v.Buffer())
	if err != nil {
		return err
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

func parseInput(username, input string) (*request.Request, error) {
	args := strings.Split(input, ">")
	msg, _ := serverRequest.NewMessageRequest(username, strings.TrimSpace(args[0]), args[1])
	req, err := msg.GenerateRequest()

	return req, err
}
