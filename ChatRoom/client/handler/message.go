package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func (c *ClientHandler) HandlePrivateMessage(body []byte) error {
	msg := &request.PrivateMessage{}
	err := json.Unmarshal(body, msg)
	if err != nil {
		logrus.Errorf("client private message handler: err while unmarshalling message body: %s", err.Error())
		return err
	}

	c.messages <- msgToString(msg)
	return nil
}

func (c *ClientHandler) Send(req *request.Request, err error) error {
	if err != nil {
		c.messages <- err.Error()
		return err
	}

	c.client.Out <- req

	return nil
}

func (c *ClientHandler) ParseInput(g *gocui.Gui, v *gocui.View) error {
	var req *request.Request
	var err error

	args := strings.Split(v.Buffer(), ">")
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}

	switch args[0] {
	case "file":
		go c.HandleWriteFile(*c.username, args[1], strings.TrimSpace(args[2]))
		break
	case "get":
		req, err := request.NewReadFileRequest(args[1]).GenerateRequest()
		c.Send(req, err)
		break
	case "new":
		newGp, err := request.NewCreateGroupRequest(args[1])
		if err != nil {
			c.Send(nil, err)
		}

		req, err := newGp.GenerateRequest()
		c.Send(req, err)
		break
	case "add gp":
		addGp, err := request.NewAddToGroupRequest(args[1], strings.TrimSpace(args[2]))
		if err != nil {
			c.Send(nil, err)
		}

		req, err := addGp.GenerateRequest()
		c.Send(req, err)
	case "rm gp":
		rmGp, err := request.NewRmFromGroupRequest(args[1], strings.TrimSpace(args[2]))
		if err != nil {
			c.Send(nil, err)
		}

		req, err := rmGp.GenerateRequest()
		c.Send(req, err)
	case "gp":
		msg, err := request.NewMsgToGroupRequest(args[1], strings.TrimSpace(args[2]))
		if err != nil {
			c.Send(nil, err)
		}

		req, err := msg.GenerateRequest()
		c.Send(req, err)
	default:
		if args[0] != "all" && !c.contains(args[0]) {
			c.Send(nil, errors.New("user does not exist"))
			return nil
		}

		msg, _ := request.NewMessageRequest(*c.username, args[0], args[1])
		req, err = msg.GenerateRequest()
		if err == nil && msg.To != "all" && req.Type == request.PrivateMessageType {
			msg.From = fmt.Sprintf("You>> %s", msg.To)
			c.messages <- msgToString(msg)
		}

		c.Send(req, err)
	}

	g.Update(func(g *gocui.Gui) error {
		v.Clear()
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)
		return nil
	})
	return err
}

func msgToString(msg *request.PrivateMessage) string {
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
