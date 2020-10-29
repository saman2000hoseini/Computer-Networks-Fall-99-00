package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"google.golang.org/protobuf/proto"
)

type ClientHandler struct {
	Client *model.Client
	db     *gorm.DB
}

func NewClientHandler(client *model.Client, db *gorm.DB) ClientHandler {
	return ClientHandler{Client: client, db: db}
}

func (c ClientHandler) StartListening() {
	for {
		var req []byte

		_, err := c.Client.Reader.Read(req)
		if err != nil {
			//TODO
		}

		handleRequest(unmarshal(req))
	}
}

func unmarshal(req []byte) request.Request {
	request := request.Request{}
	proto.Unmarshal(req, &request)

	return request
}

func handleRequest(request request.Request) {

}
