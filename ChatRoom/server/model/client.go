package model

import (
	"bufio"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg/request"
	"net"
)

type Client struct {
	Connection net.Conn
	Username   string
	Writer     *bufio.Writer
	Reader     *bufio.Reader
	Out        chan *request.Request
	In         chan *request.Request
}

func NewClient(connection net.Conn) *Client {
	return &Client{
		Connection: connection,
		Writer:     bufio.NewWriter(connection),
		Reader:     bufio.NewReader(connection),
		Out:        make(chan *request.Request, 100),
		In:         make(chan *request.Request, 100),
	}
}
