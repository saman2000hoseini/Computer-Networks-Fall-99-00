package handler

import (
	"encoding/json"
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const BufferSize = 4096

func (c *ClientHandler) HandleWriteFile(initReq *request.Request, client *model.Client) error {
	fileRequest := &request.WriteFile{}
	err := json.Unmarshal(initReq.Body, fileRequest)
	if err != nil {
		logrus.Errorf("write file handler: error while unmarshalling request: %s", err.Error())
		return err
	}

	counter := int64(0)
	path := getPath(fileRequest.FileName)

	var file *os.File
	if fileRequest.Count == counter {
		file, err = os.Create(path)
		if err != nil {
			return err
		}
	} else {
		client.In <- initReq
		return nil
	}

	defer file.Close()

	for {
		if counter*BufferSize >= fileRequest.Size {
			fmt.Println("New File Received")
			if *fileRequest.To == "all" {
				req, err := response.NewGlobalMessageResponse(newGlobalMsg(*fileRequest.From,
					generateMsg(fileRequest.FileName)), nil).GenerateResponse()

				c.informAll(req)
				return err
			}
			if c.clients[c.clientsID[*fileRequest.To]] != nil {
				req, err := response.NewMessageResponse(*fileRequest.From, *fileRequest.To,
					generateMsg(fileRequest.FileName)).GenerateResponse()

				c.clients[c.clientsID[*fileRequest.To]].Out <- req
				return err
			}
			args := strings.Split(*fileRequest.To, ">")
			if args[0] == "gp" {
				msg := &request.MsgToGroup{
					GroupName: args[1],
					Msg:       generateMsg(fileRequest.FileName),
				}

				req, err := msg.GenerateRequest()
				go c.HandleMsgToGroup(req.Body, client)

				return err
			}

			req, err := response.NewMessageResponse(*fileRequest.From, "",
				pkg.ErrUserNotFound).GenerateResponse()

			client.Out <- req
			return err
		}

		req := <-client.In
		if req.Type != request.WriteFileType {
			if req.Type == request.PrivateMessageType {
				go c.HandlePrivateMessage(req.Body, client)
			} else if req.Type == request.MsgToGroupType {
				go c.HandleMsgToGroup(req.Body, client)
			} else {
				client.In <- req
			}
			continue
		}

		fileReq := &request.WriteFile{}
		err := json.Unmarshal(req.Body, fileReq)
		if err != nil {
			logrus.Errorf("write file handler: error while unmarshalling request: %s", err.Error())
			continue
		}

		if fileReq.FileName != fileRequest.FileName {
			client.In <- req
			continue
		}

		counter++
		if fileReq.Count != counter {
			client.In <- req
			continue
		}

		file.Write(fileReq.File[:fileReq.Size])
	}
}

func getPath(filename string) string {
	return fmt.Sprintf("./server/storage/%s", filename)
}

func generateMsg(name string) string {
	return fmt.Sprintf("sent you a file: %s", name)
}
