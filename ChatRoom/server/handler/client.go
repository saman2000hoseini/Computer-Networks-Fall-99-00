package handler

import "github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"

type ClientHandler struct {
	Client model.Client
}

func NewClientHandler(client model.Client) ClientHandler {
	return ClientHandler{Client: client}
}

func (C ClientHandler) StartListening() {

}
