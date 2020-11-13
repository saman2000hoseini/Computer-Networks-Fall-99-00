package handler

import (
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"io"
	"os"
	"strings"
	"time"
)

const BufferSize = 4096

func (c *ClientHandler) HandleWriteFile(from, to, src string) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		c.Send(nil, err)
		return
	}

	if !sourceFileStat.Mode().IsRegular() {
		c.Send(nil, fmt.Errorf("%s is not a regular file", src))
		return
	}

	source, err := os.Open(src)
	if err != nil {
		c.Send(nil, err)
		return
	}
	defer source.Close()

	fName := getFileName(src)
	fSize := sourceFileStat.Size()
	count := int64(0)

	var fileReq *request.File
	fileReq, _ = request.NewFileRequest(&from, &to, fName, nil, count, fSize)
	req, err := fileReq.GenerateRequest()

	c.Send(req, nil)
	count++

	for {
		buf := make([]byte, BufferSize)
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			c.Send(nil, err)
			return
		}
		if n == 0 {
			break
		}

		fileReq, _ = request.NewFileRequest(nil, nil, fName, buf, count, int64(n))
		req, err := fileReq.GenerateRequest()

		c.Send(req, nil)
		count++
	}
}

func getFileName(filename string) string {
	path := strings.Split(filename, "/")
	return fmt.Sprintf("%d_%s", time.Now().Unix(), path[len(path)-1])
}
