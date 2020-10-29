package model

import (
	"bufio"
	"net"
)

type Client struct {
	Connection net.Conn
	Writer     *bufio.Writer
	Reader     *bufio.Reader
}

func NewClient(connection net.Conn) *Client {
	return &Client{Connection: connection, Writer: bufio.NewWriter(connection), Reader: bufio.NewReader(connection)}
}
