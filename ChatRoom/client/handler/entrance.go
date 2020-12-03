package handler

import (
	"github.com/jroimartin/gocui"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"strings"
)

func (c *ClientHandler) entrance(g *gocui.Gui, v *gocui.View) error {
	args := strings.Split(v.Buffer(), ",")
	g.Update(func(g *gocui.Gui) error {
		v.Clear()
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)
		return nil
	})

	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}

	if len(args) < 2 || len(args[0]) < 6 || len(args[1]) < 8 {
		return nil
	}

	if len(args) == 3 {
		c.username = &args[0]
		si, _ := request.NewSignUpRequest(args[0], args[1], args[2])
		req, _ := si.GenerateRequest()
		c.client.Out <- req
	} else {
		c.username = &args[0]
		si, _ := request.NewSignInRequest(args[0], args[1])
		req, _ := si.GenerateRequest()
		c.client.Out <- req
	}

	<-c.waiter

	if c.signedIn {
		g.SetViewOnTop("messages")
		g.SetViewOnTop("users")
		g.SetViewOnTop("input")
		g.SetCurrentView("input")
	}

	return nil
}
