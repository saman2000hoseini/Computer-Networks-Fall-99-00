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
			if c.clients[*fileRequest.To] != nil {
				req, err := response.NewMessageResponse(*fileRequest.From, *fileRequest.To,
					generateMsg(fileRequest.FileName)).GenerateResponse()

				c.clients[*fileRequest.To].Out <- req
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
