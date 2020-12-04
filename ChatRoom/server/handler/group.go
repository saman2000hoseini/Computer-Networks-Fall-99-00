package handler

import (
	"encoding/json"
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/pkg"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/response"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"github.com/sirupsen/logrus"
	"strconv"
)

func (c *ClientHandler) HandleCreateGroup(body []byte, client *model.Client) error {
	info := &request.CreateGroup{}
	err := json.Unmarshal(body, info)
	if err != nil {
		logrus.Errorf("group create handler: erro while unmarshalling request: %s", err.Error())
		return err
	}

	group := model.NewGroup(info.Name, client.ID)

	err = c.groupRepo.Save(group)
	if err == nil {
		req, err := response.NewMessageResponse("server", client.Username,
			"group created: "+info.Name).GenerateResponse()

		client.Out <- req
		return err
	}

	req, err := response.NewMessageResponse("server", client.Username,
		"error creating group: "+err.Error()).GenerateResponse()
	client.Out <- req

	return err
}

func (c *ClientHandler) HandleAddToGroup(body []byte, client *model.Client) error {
	info := &request.AddToGroup{}
	err := json.Unmarshal(body, info)
	if err != nil {
		logrus.Errorf("sign up handler: err while unmarshalling request: %s", err.Error())
		return err
	}

	group, err := c.groupRepo.Find(info.GroupName)
	if err != nil {
		req, err := response.NewMessageResponse("server", client.Username,
			"error finding group: "+err.Error()).GenerateResponse()
		client.Out <- req

		return err
	}

	if isMember(c.clientsID[client.Username], group) == -1 {
		req, err := response.NewMessageResponse("server", client.Username,
			pkg.ErrNoAccess).GenerateResponse()
		client.Out <- req

		return err
	}

	if isMember(c.clientsID[info.Username], group) != -1 {
		req, err := response.NewMessageResponse("server", client.Username,
			"user already exists").GenerateResponse()
		client.Out <- req

		return err
	}

	group.Members = append(group.Members, info.Username)
	err = c.groupRepo.Update(group)
	if err != nil {
		return err
	}

	req, err := response.NewMessageResponse(groupMsg(group.Name), "all",
		client.Username+" added new user to group: "+info.Username).GenerateResponse()
	c.sendMsg(group, req)

	return err
}

func (c *ClientHandler) HandleMsgToGroup(body []byte, client *model.Client) error {
	Msg := &request.MsgToGroup{}
	err := json.Unmarshal(body, Msg)
	if err != nil {
		logrus.Errorf("sign up handler: err while unmarshalling request: %s", err.Error())
		return err
	}

	group, err := c.groupRepo.Find(Msg.GroupName)
	if err != nil {
		req, err := response.NewMessageResponse("server", client.Username,
			"error finding group: "+err.Error()).GenerateResponse()
		client.Out <- req

		return err
	}

	if isMember(c.clientsID[client.Username], group) == -1 {
		req, err := response.NewMessageResponse("server", client.Username,
			pkg.ErrNoAccess).GenerateResponse()
		client.Out <- req

		return err
	}

	req, err := response.NewMessageResponse(groupMsg(group.Name), "all",
		client.Username+": "+Msg.Msg).GenerateResponse()
	c.sendMsg(group, req)

	return err
}

func (c *ClientHandler) HandleRmFromGroup(body []byte, client *model.Client) error {
	info := &request.RmFromGroup{}
	err := json.Unmarshal(body, info)
	if err != nil {
		logrus.Errorf("sign up handler: err while unmarshalling request: %s", err.Error())
		return err
	}

	group, err := c.groupRepo.Find(info.GroupName)
	if err != nil {
		req, err := response.NewMessageResponse("server", client.Username,
			"error finding group: "+err.Error()).GenerateResponse()
		client.Out <- req

		return err
	}

	if client.ID != group.Admin {
		req, err := response.NewMessageResponse("server", client.Username,
			pkg.ErrNoAccess).GenerateResponse()
		client.Out <- req

		return err
	}

	index := isMember(c.clientsID[info.Username], group)
	if index == -1 {
		req, err := response.NewMessageResponse("server", client.Username,
			"user is not in the group").GenerateResponse()
		client.Out <- req

		return err
	}

	group.Members[index] = group.Members[len(group.Members)-1]
	group.Members[len(group.Members)-1] = ""
	group.Members = group.Members[:len(group.Members)-1]
	err = c.groupRepo.Update(group)

	req, err := response.NewMessageResponse(groupMsg(group.Name), "all",
		c.clientsUser[group.Admin]+" removed user from group: "+info.Username).GenerateResponse()
	c.sendMsg(group, req)
	if c.clients[c.clientsID[info.Username]] != nil {
		c.clients[c.clientsID[info.Username]].Out <- req
	}

	return err
}

func isMember(id uint64, group model.Group) int {
	for i := range group.Members {
		if group.Members[i] == strconv.FormatUint(id, 10) {
			return i
		}
	}

	return -1
}

func groupMsg(name string) string {
	return fmt.Sprintf("\u001B[1;35m%s\u001B", name)
}

func (c *ClientHandler) sendMsg(group model.Group, req *request.Request) {
	for i := range group.Members {
		if c.clients[c.clientsID[group.Members[i]]] != nil {
			c.clients[c.clientsID[group.Members[i]]].Out <- req
		}
	}
}
