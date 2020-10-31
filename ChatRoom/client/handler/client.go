package handler

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/client/view"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	serverRequest "github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"strings"
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
		messages:    make(chan string, 20),
		Users:       make([]string, 0),
		usersChange: make(chan bool, 5),
	}
}

func (c *ClientHandler) Handle() {
	go c.StartListening()
	go c.Request()

	c.entrance()

	c.gui, _ = gocui.NewGui(gocui.OutputNormal)
	defer c.gui.Close()

	c.gui.SetManagerFunc(view.Layout)
	c.gui.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, c.Send)
	c.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, c.Disconnect)
	c.gui.MainLoop()
	go c.writeMessage()
	go c.updateUsers()
	c.waiter <- true
}

func (c *ClientHandler) StartListening() {
	for {
		req := make([]byte, 1024)

		count, err := c.client.Reader.Read(req)
		if err != nil {
			//TODO
		}

		protoRequest, err := unmarshal(req[:count])
		if err != nil {
			logrus.Errorf("client handler: err while unmarshalling proto: %s", err.Error())
			continue
		}

		go c.handleRequest(protoRequest)
	}
}

func unmarshal(req []byte) (*request.Request, error) {
	protoRequest := &request.Request{}
	err := proto.Unmarshal(req, protoRequest)
	if err != nil {
		return nil, err
	}

	return protoRequest, nil
}

func (c *ClientHandler) handleRequest(req *request.Request) {
	var err error

	switch req.Type {
	case response.SignType:
		err = c.HandleSign(req.Body)
		break
	case response.PrivateMessageType:
		err = c.HandlePrivateMessage(req.Body)
		break
	case response.GlobalMessageType:
		err = c.HandleGlobalMessage(req.Body)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}

func printFirstMenu() {
	fmt.Printf("1) Sign up\n2) Sign in\n")
}

func (c *ClientHandler) Request() {
	for {
		req := <-c.client.Out

		out, err := proto.Marshal(req)
		if err != nil {
			logrus.Errorf("client: error while marshalling request proto: %s", err.Error())
			return
		}

		_, err = c.client.Writer.Write(out)
		if err != nil {
			logrus.Errorf("client: error while writing request proto: %s", err.Error())
			return
		}

		err = c.client.Writer.Flush()
		if err != nil {
			logrus.Errorf("client: error while flushing writer: %s", err.Error())
			return
		}
	}
}

func (c *ClientHandler) entrance() {
	for !c.signedIn {
		printFirstMenu()
		var cmd int
		fmt.Scanf("%d\n", &cmd)
		if cmd == 1 {
			var username, password, email string
			fmt.Scanf("%s\n%s\n%s\n", &username, &password, &email)
			c.username = &username
			su, _ := serverRequest.NewSignUpRequest(username, password, email)
			req, _ := su.GenerateRequest()
			c.client.Out <- req
		} else if cmd == 2 {
			var username, password string
			fmt.Scanf("%s\n%s\n", &username, &password)
			c.username = &username
			si, _ := serverRequest.NewSignInRequest(username, password)
			req, _ := si.GenerateRequest()
			c.client.Out <- req
		}

		<-c.waiter
	}
	println("finished")
}

func (c *ClientHandler) Disconnect(g *gocui.Gui, v *gocui.View) error {
	c.client.Connection.Close()
	return gocui.ErrQuit
}

func (c *ClientHandler) writeMessage() {
	messagesView, _ := c.gui.View("messages")

	for {
		msg := <-c.messages

		c.gui.Update(func(g *gocui.Gui) error {
			fmt.Fprintln(messagesView, msg)
			return nil
		})
	}
}

func (c *ClientHandler) updateUsers() {
	usersView, _ := c.gui.View("users")

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
