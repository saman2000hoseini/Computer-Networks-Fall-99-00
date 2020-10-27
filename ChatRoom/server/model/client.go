package model

import (
	"bufio"
	"net"
)

type Client struct {
	connection net.Conn
	writer     *bufio.Writer
	reader     *bufio.Reader
}

func NewClient(connection net.Conn) *Client {
	return &Client{connection: connection, writer: bufio.NewWriter(connection), reader: bufio.NewReader(connection)}
}
