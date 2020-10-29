package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/utils"
	serverRequest "github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type ClientHandler struct {
	db      *gorm.DB
	clients map[string]*model.Client
}

func NewClientHandler(db *gorm.DB) ClientHandler {
	return ClientHandler{db: db}
}

func (c ClientHandler) StartListening(client *model.Client) {
	for {
		var req []byte

		_, err := client.Reader.Read(req)
		if err != nil {
			//TODO
		}

		protoRequest, err := unmarshal(req)
		if err != nil {
			logrus.Errorf("client handler: err while unmarshalling proto: %s", err.Error())
			continue
		}

		go c.handleRequest(protoRequest, client)
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

func (c ClientHandler) handleRequest(req *request.Request, client *model.Client) {
	resp := &request.Request{}
	var err error

	switch req.Type {
	case serverRequest.SignInType:
		resp, err = c.HandleSignIn(req.Body, client)
		break
	case serverRequest.SignUpType:
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

	out, err := proto.Marshal(resp)
	if err != nil {
		logrus.Errorf("client handler: err while marshalling proto: %s", err.Error())
		return
	}

	_, err = client.Writer.Write(out)
	if err != nil {
		logrus.Errorf("client handler: err while writing proto: %s", err.Error())
		return
	}

	err = client.Writer.Flush()
	if err != nil {
		logrus.Errorf("client handler: err while flushing writer: %s", err.Error())
		return
	}
}
