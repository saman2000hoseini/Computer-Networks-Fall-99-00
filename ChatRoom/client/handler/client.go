package handler

import (
	"encoding/json"
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/client/view"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"
	"time"
)

type ClientHandler struct {
	client      *model.Client
	username    *string
	signedIn    bool
	waiter      chan bool
	gui         *gocui.Gui
	messages    chan string
	usersChange chan bool
	Users       []string
}

func NewClientHandler(client *model.Client) *ClientHandler {
	return &ClientHandler{
		client:      client,
		signedIn:    false,
		waiter:      make(chan bool),
		messages:    make(chan string, 100),
		Users:       make([]string, 0),
		usersChange: make(chan bool, 5),
	}
}

func (c *ClientHandler) Handle() {
	var err error
	go c.StartListening()
	go c.Request()
	go c.handleRequest()

	c.gui, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	c.gui.SetManagerFunc(view.Layout)
	c.gui.SetKeybinding("main-menu", gocui.KeyEnter, gocui.ModNone, c.entrance)
	c.gui.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, c.ParseInput)
	c.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, c.Disconnect)
	go c.writeMessage()
	go c.updateUsers()
	c.gui.MainLoop()
}

func (c *ClientHandler) StartListening() {
	defer c.gui.Close()

	for {
		req, err := c.client.Reader.ReadString('\n')
		if err != nil {
			//TODO
			os.Exit(1)
		}

		jsonRequest, err := unmarshal([]byte(req))
		if err != nil {
			logrus.Errorf("client handler: err while unmarshalling: %s", err.Error())
			continue
		}

		c.client.In <- jsonRequest
	}
}

func unmarshal(req []byte) (*request.Request, error) {
	jsonRequest := &request.Request{}
	err := json.Unmarshal(req, jsonRequest)
	if err != nil {
		return nil, err
	}

	return jsonRequest, nil
}

func (c *ClientHandler) handleRequest() {
	for {
		req := <-c.client.In
		var err error

		switch req.Type {
		case response.ChangeUserType:
			err = c.updateUsername(req.Body)
		case response.SignType:
			err = c.HandleSign(req.Body)
		case response.PrivateMessageType:
			err = c.HandlePrivateMessage(req.Body)
		case response.GlobalMessageType:
			err = c.HandleGlobalMessage(req.Body)
		case response.DownloadFileType:
			err = c.HandleGetFile(req)
		}
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (c *ClientHandler) Request() {
	for {
		req := <-c.client.Out

		out, err := json.Marshal(req)
		if err != nil {
			logrus.Errorf("client: error while marshalling request: %s", err.Error())
			return
		}

		_, err = c.client.Writer.WriteString(string(out) + "\n")
		if err != nil {
			logrus.Errorf("client: error while writing request: %s", err.Error())
			return
		}

		err = c.client.Writer.Flush()
		if err != nil {
			logrus.Errorf("client: error while flushing writer: %s", err.Error())
			return
		}
	}
}

func (c *ClientHandler) Disconnect(g *gocui.Gui, v *gocui.View) error {
	c.client.Connection.Close()
	return gocui.ErrQuit
}

func (c *ClientHandler) updateUsername(body []byte) error {
	res := &response.ChangeUser{}
	err := json.Unmarshal(body, res)
	if err != nil {
		logrus.Errorf("client private message handler: err while unmarshalling message body: %s", err.Error())
		return err
	}

	c.username = &res.Username
	return nil
}

func (c *ClientHandler) writeMessage() {
	<-time.Tick(500 * time.Millisecond)
	messagesView, err := c.gui.View("messages")
	if err != nil {
		panic(err)
	}

	for {
		msg := <-c.messages

		c.gui.Update(func(g *gocui.Gui) error {
			fmt.Fprintln(messagesView, msg)
			return nil
		})
	}
}

func (c *ClientHandler) updateUsers() {
	<-time.Tick(1 * time.Second)
	usersView, err := c.gui.View("users")
	if err != nil {
		panic(err)
	}

	for {
		<-c.usersChange

		c.gui.Update(func(g *gocui.Gui) error {
			usersView.Title = fmt.Sprintf(" %d users: ", len(c.Users))
			usersView.Clear()
			fmt.Fprint(usersView, strings.Join(c.Users, "\n"))
			return nil
		})
	}
}
