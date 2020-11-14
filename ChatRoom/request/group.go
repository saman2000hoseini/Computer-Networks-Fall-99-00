package request

import (
	"encoding/json"
	"errors"
)

const (
	CreateGroupType = "create-group"
	AddToGroupType  = "add-to-group"
	MsgToGroupType  = "msg-to-group"
	RmFromGroupType = "rm-from-group"
)

type CreateGroup struct {
	Name string `json:"name"`
}

type AddToGroup struct {
	GroupName string `json:"group_name"`
	Username  string `json:"username"`
}

type MsgToGroup struct {
	GroupName string `json:"group_name"`
	Msg       string `json:"msg"`
}

type RmFromGroup struct {
	GroupName string `json:"group_name"`
	Username  string `json:"username"`
}

func NewCreateGroupRequest(name string) (*CreateGroup, error) {
	if len(name) < 6 {
		return nil, errors.New("name is not under 6 chars")
	}

	return &CreateGroup{
		Name: name,
	}, nil
}

func NewAddToGroupRequest(gName, username string) (*AddToGroup, error) {
	if len(gName) < 6 || len(username) < 6 {
		return nil, errors.New("name is not under 6 chars")
	}

	return &AddToGroup{
		GroupName: gName,
		Username:  username,
	}, nil
}

func NewMsgToGroupRequest(gName, msg string) (*MsgToGroup, error) {
	if len(gName) < 6 {
		return nil, errors.New("name is not under 6 chars")
	}

	return &MsgToGroup{
		GroupName: gName,
		Msg:       msg,
	}, nil
}

func NewRmFromGroupRequest(gName, username string) (*RmFromGroup, error) {
	if len(gName) < 6 || len(username) < 6 {
		return nil, errors.New("name is not under 6 chars")
	}

	return &RmFromGroup{
		GroupName: gName,
		Username:  username,
	}, nil
}

func (g CreateGroup) GenerateRequest() (*Request, error) {
	body, err := json.Marshal(g)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(CreateGroupType, body), nil
}

func (g AddToGroup) GenerateRequest() (*Request, error) {
	body, err := json.Marshal(g)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(AddToGroupType, body), nil
}

func (g MsgToGroup) GenerateRequest() (*Request, error) {
	body, err := json.Marshal(g)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(MsgToGroupType, body), nil
}

func (g RmFromGroup) GenerateRequest() (*Request, error) {
	body, err := json.Marshal(g)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(RmFromGroupType, body), nil
}
