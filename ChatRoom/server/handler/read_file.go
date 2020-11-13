package handler

import (
	"encoding/json"
	"errors"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func (c *ClientHandler) HandleReadFile(body []byte, client *model.Client) error {
	fileRequest := &request.ReadFile{}
	err := json.Unmarshal(body, fileRequest)
	if err != nil {
		logrus.Errorf("sign in handler: err while unmarshalling request: %s", err.Error())
		return err
	}

	path := getPath(fileRequest.FileName)
	sourceFileStat, err := os.Stat(path)
	if err != nil {
		req, err := response.NewMessageResponse("Server", client.Username, pkg.ErrFileNotFound).GenerateResponse()
		client.Out <- req

		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return errors.New(pkg.ErrInternal)
	}

	source, err := os.Open(path)
	if err != nil {
		return errors.New(pkg.ErrInternal)
	}

	defer source.Close()

	fSize := sourceFileStat.Size()
	count := int64(0)

	var fileReq *response.DownloadFile
	fileReq = response.NewDownloadFileResponse(fileRequest.FileName, nil, count, fSize)
	req, err := fileReq.GenerateResponse()

	client.Out <- req
	count++

	for {
		buf := make([]byte, BufferSize)
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return errors.New(pkg.ErrInternal)
		}
		if n == 0 {
			break
		}

		res, err := response.NewDownloadFileResponse(fileRequest.FileName, buf, count, int64(n)).GenerateResponse()

		client.Out <- res
		count++
	}

	return nil
}
