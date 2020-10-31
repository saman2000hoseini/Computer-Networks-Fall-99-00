package handler

import (
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/utils"
	serverRequest "github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

type ClientHandler struct {
	db        *gorm.DB
	clients   map[string]*model.Client
	clientIDs []string
}

func NewClientHandler(db *gorm.DB) *ClientHandler {
	return &ClientHandler{db: db, clients: make(map[string]*model.Client), clientIDs: make([]string, 0)}
}

func (c *ClientHandler) StartListening(client *model.Client) {
	for {
		req := make([]byte, 1024)

		count, err := client.Reader.Read(req)
		if err != nil {
			delete(c.clients, client.Username)
			response.LogOut(&c.clientIDs, client.Username)
			c.informJoin(client.Username, false)
			break
		}
		fmt.Println("request received")

		protoRequest, err := unmarshal(req[:count])
		if err != nil {
			logrus.Errorf("client handler: err while unmarshalling proto: %s", err.Error())
			continue
		}

		go c.handleRequest(protoRequest, client)
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

func (c *ClientHandler) handleRequest(req *request.Request, client *model.Client) {
	resp := &request.Request{}
	var err error

	switch req.Type {
	case serverRequest.SignInType:
		resp, err = c.HandleSignIn(req.Body, client)
		break
	case serverRequest.SignUpType:
		fmt.Println("signup request")
		resp, err = c.HandleSignUp(req.Body, client)
		break
	case serverRequest.PrivateMessageType:
		client, resp, err = c.HandlePrivateMessage(req.Body, client)
		break

	}

	if err != nil {
		resp.Type = utils.ErrInternal

		out, err := proto.Marshal(resp)
		if err != nil {
			logrus.Errorf("client handler: err while marshalling error proto: %s", err.Error())
			return
		}

		_, err = client.Writer.Write(out)
		if err != nil {
			logrus.Errorf("client handler: err while writing error proto: %s", err.Error())
			return
		}

		err = client.Writer.Flush()
		if err != nil {
			logrus.Errorf("client handler: err while flushing writer: %s", err.Error())
			return
		}

		return
	}

	client.Out <- resp
}

func (c *ClientHandler) Respond(client *model.Client) {
	for {
		resp := <-client.Out
		out, err := proto.Marshal(resp)
		if err != nil {
			logrus.Errorf("client handler: error while marshalling proto: %s", err.Error())
			return
		}

		_, err = client.Writer.Write(out)
		if err != nil {
			logrus.Errorf("client handler: error while writing proto: %s", err.Error())
			return
		}

		err = client.Writer.Flush()
		if err != nil {
			logrus.Errorf("client handler: error while flushing writer: %s", err.Error())
			return
		}
	}
}

func (c *ClientHandler) informJoin(username string, joined bool) {
	gm, err := response.NewGlobalMessageResponse(username, &joined).GenerateResponse()
	if err != nil {
		logrus.Errorf("client handler: error while generating global message: %s", err.Error())
		return
	}
	fmt.Println(gm)
	fmt.Println(c.clients)
	fmt.Println(c.clientIDs)
	c.informAll(gm)
}

func (c *ClientHandler) informAll(req *request.Request) {
	for _, client := range c.clients {
		client.Out <- req
	}
}
