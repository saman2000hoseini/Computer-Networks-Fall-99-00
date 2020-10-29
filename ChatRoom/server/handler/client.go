package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	request2 "github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type ClientHandler struct {
	db      *gorm.DB
	clients map[uint64]*model.Client
}

func NewClientHandler(db *gorm.DB) ClientHandler {
	return ClientHandler{db: db}
}

func (c ClientHandler) StartListening(Client *model.Client) {
	for {
		var req []byte

		_, err := Client.Reader.Read(req)
		if err != nil {
			//TODO
		}

		protoRequest, err := unmarshal(req)
		if err != nil {
			logrus.Errorf("client handler: err while unmarshalling proto: %s", err.Error())
			continue
		}

		c.handleRequest(protoRequest)
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

func (c ClientHandler) handleRequest(request *request.Request) {
	switch request.Type {
	case request2.SignInType:
		c.HandleSignIn(request.Body)
		break
	case request2.SignUpType:
		c.HandleSignUp(request.Body)
		break
	case request2.PrivateMessageType:
		c.HandlePrivateMessage(request.Body)
		break

	}
}
