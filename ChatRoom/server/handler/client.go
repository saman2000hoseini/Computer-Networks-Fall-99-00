package handler

import (
	"encoding/json"
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ClientHandler struct {
	userRepo    model.UserRepo
	groupRepo   model.GroupRepo
	clients     map[uint64]*model.Client
	clientIDs   []uint64
	clientsUser map[uint64]string
	clientsID   map[string]uint64
}

func NewClientHandler(db *gorm.DB) *ClientHandler {
	return &ClientHandler{userRepo: model.SQLUserRepo{DB: db}, groupRepo: model.SQLGroupRepo{DB: db},
		clients: make(map[uint64]*model.Client), clientIDs: make([]uint64, 0), clientsUser: make(map[uint64]string),
		clientsID: make(map[string]uint64)}
}

func (c *ClientHandler) StartListening(client *model.Client) {
	for {
		req, err := client.Reader.ReadString('\n')
		if err != nil {
			delete(c.clients, client.ID)
			delete(c.clientsID, client.Username)
			delete(c.clientsUser, client.ID)
			response.LogOut(&c.clientIDs, client.ID)
			c.informJoin(client.Username, false)
			break
		}

		jsonRequest, err := unmarshal([]byte(req))
		if err != nil {
			logrus.Errorf("client handler: err while unmarshalling proto: %s", err.Error())
			continue
		}

		client.In <- jsonRequest
	}
}

func unmarshal(req []byte) (*request.Request, error) {
	jsonReq := &request.Request{}
	err := json.Unmarshal(req, jsonReq)
	if err != nil {
		return nil, err
	}

	return jsonReq, nil
}

func (c *ClientHandler) HandleRequest(client *model.Client) {
	for {
		req := <-client.In
		resp := &request.Request{}
		var err error

		fmt.Println(req.Type + " request")
		switch req.Type {
		case request.ChangeInfoType:
			err = c.HandleChangeInfo(req.Body, client)
			break
		case request.SignInType:
			err = c.HandleSignIn(req.Body, client)
			break
		case request.SignUpType:
			err = c.HandleSignUp(req.Body, client)
			break
		case request.PrivateMessageType:
			err = c.HandlePrivateMessage(req.Body, client)
			break
		case request.WriteFileType:
			err = c.HandleWriteFile(req, client)
			break
		case request.ReadFileType:
			err = c.HandleReadFile(req.Body, client)
			break
		case request.CreateGroupType:
			err = c.HandleCreateGroup(req.Body, client)
			break
		case request.AddToGroupType:
			err = c.HandleAddToGroup(req.Body, client)
			break
		case request.MsgToGroupType:
			err = c.HandleMsgToGroup(req.Body, client)
			break
		case request.RmFromGroupType:
			err = c.HandleRmFromGroup(req.Body, client)
			break
		}

		if err != nil {
			fmt.Println(err.Error())
			if resp == nil {
				resp = &request.Request{}
			}

			resp.Type = pkg.ErrInternal

			out, err := json.Marshal(resp)
			if err != nil {
				logrus.Errorf("client handler: err while marshalling error proto: %s", err.Error())
				return
			}

			_, err = client.Writer.WriteString(string(out) + "\n")
			if err != nil {
				logrus.Errorf("client handler: err while writing error proto: %s", err.Error())
				return
			}

			err = client.Writer.Flush()
			if err != nil {
				logrus.Errorf("client handler: err while flushing writer: %s", err.Error())
				return
			}
		}
	}
}

func (c *ClientHandler) Respond(client *model.Client) {
	for {
		resp := <-client.Out
		out, err := json.Marshal(resp)
		if err != nil {
			logrus.Errorf("client handler: error while marshalling proto: %s", err.Error())
			return
		}

		_, err = client.Writer.WriteString(string(out) + "\n")
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
	c.informAll(gm)
}

func (c *ClientHandler) informAll(req *request.Request) {
	for _, client := range c.clients {
		client.Out <- req
	}
}
