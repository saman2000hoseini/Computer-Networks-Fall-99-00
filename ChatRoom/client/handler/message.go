package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg"
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
	if len(args) == 1 {
		return nil
	}

	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}

	switch args[0] {
	case "change":
		args := strings.Split(args[1], ",")
		if len(args) != 3 {
			return nil
		}

		for i := range args {
			args[i] = strings.TrimSpace(args[i])
		}

		chInfo, err := request.NewChangeInfoRequest(args[0], args[1], args[2])
		if err != nil {
			c.Send(nil, err)
		}

		req, err := chInfo.GenerateRequest()
		c.Send(req, err)
	case "file":
		if args[1] == "gp" {
			if len(args) < 4 {
				return nil
			}

			go c.HandleWriteFile(*c.username, args[1]+">"+args[2], strings.TrimSpace(args[3]))
		} else {
			if len(args) < 3 {
				return nil
			}

			go c.HandleWriteFile(*c.username, args[1], strings.TrimSpace(args[2]))
		}
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
		if len(args) < 3 {
			return nil
		}

		addGp, err := request.NewAddToGroupRequest(args[1], strings.TrimSpace(args[2]))
		if err != nil {
			c.Send(nil, err)
		}

		req, err := addGp.GenerateRequest()
		c.Send(req, err)
	case "rm gp":
		if len(args) < 3 {
			return nil
		}

		rmGp, err := request.NewRmFromGroupRequest(args[1], strings.TrimSpace(args[2]))
		if err != nil {
			c.Send(nil, err)
		}

		req, err := rmGp.GenerateRequest()
		c.Send(req, err)
	case "gp":
		if len(args) < 3 {
			return nil
		}

		msg, err := request.NewMsgToGroupRequest(args[1], strings.TrimSpace(args[2]))
		if err != nil {
			c.Send(nil, err)
		}

		req, err := msg.GenerateRequest()
		c.Send(req, err)
	default:
		if args[0] != "all" && !c.contains(args[0]) {
			c.Send(nil, errors.New(pkg.ErrUserDoesntExist))
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
