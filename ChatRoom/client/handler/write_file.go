package handler

import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func (c *ClientHandler) HandleGetFile(initReq *request.Request) error {
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
		c.client.In <- initReq
		return nil
	}

	defer file.Close()

	for {
		if counter*BufferSize >= fileRequest.Size {
			c.messages <- fmt.Sprintf("[%v] \u001B[1;33m downloading %s finished \u001B[0m",
				time.Now().Local(), fileRequest.FileName)

			return nil
		}

		req := <-c.client.In
		if req.Type != response.DownloadFileType {
			if req.Type == response.PrivateMessageType {
				go c.HandlePrivateMessage(req.Body)
			} else if req.Type == response.GlobalMessageType {
				go c.HandleGlobalMessage(req.Body)
			} else {
				c.client.In <- req
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
			c.client.In <- req
			continue
		}

		counter++
		if fileReq.Count != counter {
			c.client.In <- req
			continue
		}

		file.Write(fileReq.File[:fileReq.Size])
	}
}

func getPath(filename string) string {
	i := strings.Index(filename, "_")
	fName := filename[i+1:]
	return fmt.Sprintf("./client/downloads/%s", fName)
}
