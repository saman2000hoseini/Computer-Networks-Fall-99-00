package request

import (
	"encoding/json"
	"errors"
)

const (
	WriteFileType = "write-file"
	ReadFileType  = "read-file"
)

type WriteFile struct {
	From     *string `json:"from"`
	To       *string `json:"to"`
	FileName string  `json:"file_name"`
	Count    int64   `json:"count"`
	Size     int64   `json:"size"`
	File     []byte  `json:"file"`
}

type ReadFile struct {
	FileName string `json:"file_name"`
}

func NewWriteFileRequest(from, to *string, fileName string, file []byte, count int64, size int64) (*WriteFile, error) {
	return &WriteFile{
		From:     from,
		To:       to,
		FileName: fileName,
		Count:    count,
		Size:     size,
		File:     file,
	}, nil
}

func (f WriteFile) GenerateRequest() (*Request, error) {
	body, err := json.Marshal(f)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(WriteFileType, body), nil
}

func NewReadFileRequest(fileName string) *ReadFile {
	return &ReadFile{FileName: fileName}
}

func (f ReadFile) GenerateRequest() (*Request, error) {
	body, err := json.Marshal(f)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(ReadFileType, body), nil
}
