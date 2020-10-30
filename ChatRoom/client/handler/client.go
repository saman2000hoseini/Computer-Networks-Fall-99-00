package handler

import (
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	serverRequest "github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type ClientHandler struct {
	client   *model.Client
	username *string
	signedIn bool
	waiter   chan bool
}

func NewClientHandler(client *model.Client) ClientHandler {
	return ClientHandler{client: client, signedIn: false}
}

func (c ClientHandler) Handle() {
	go c.StartListening()

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
			c.Request(req)
		} else if cmd == 2 {
			var username, password string
			fmt.Scanf("%s\n%s\n", &username, &password)
			c.username = &username
			si, _ := serverRequest.NewSignInRequest(username, password)
			req, _ := si.GenerateRequest()
			c.Request(req)
		}

		<-c.waiter
	}

	for {
		printMenu()
		var cmd int
		fmt.Scanf("%d\n", &cmd)
		var username, msg string
		fmt.Scanf("%s\n%s\n%s\n", &username, &msg)
		su, _ := serverRequest.NewMessageRequest(*c.username, username, msg)
		req, _ := su.GenerateRequest()
		c.Request(req)
	}
}

func (c ClientHandler) StartListening() {
	for {
		req := make([]byte, 1024)

		count, err := c.client.Reader.Read(req)
		if err != nil {
			//TODO
		}
		fmt.Println("response received")

		protoRequest, err := unmarshal(req[:count])
		if err != nil {
			logrus.Errorf("client handler: err while unmarshalling proto: %s", err.Error())
			continue
		}

		go c.handleRequest(protoRequest)
	}
}

func unmarshal(req []byte) (*request.Request, error) {
	protoRequest := request.Request{}
	err := proto.Unmarshal(req, &protoRequest)
	if err != nil {
		return nil, err
	}

	return &protoRequest, nil
}

func (c ClientHandler) handleRequest(req *request.Request) {
	var err error

	switch req.Type {
	case serverRequest.SignInType:
		err = c.HandleSign(req.Body)
		break
	case serverRequest.SignUpType:
		err = c.HandleSign(req.Body)
		break
	case serverRequest.PrivateMessageType:
		err = c.HandlePrivateMessage(req.Body)
		break

	}
	if err != nil {
		fmt.Println(err.Error())
	}
}

func printFirstMenu() {
	fmt.Printf("1) Sign up\n2) Sign in\n")
}

func printMenu() {
	fmt.Printf("1) Private message\n")
}

func (c *ClientHandler) Request(req *request.Request) {
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
