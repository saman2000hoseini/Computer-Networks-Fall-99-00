package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
)

type ClientHandler struct {
	Client *model.Client
	db     *gorm.DB
}

func NewClientHandler(client *model.Client, db *gorm.DB) ClientHandler {
	return ClientHandler{Client: client, db: db}
}

func (C ClientHandler) StartListening() {

}
