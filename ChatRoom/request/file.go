package request

import (
	"encoding/json"
	"errors"
)

const FileType = "file"

type File struct {
	From     *string `json:"from"`
	To       *string `json:"to"`
	FileName string  `json:"file_name"`
	Count    int64   `json:"count"`
	Size     int64   `json:"size"`
	File     []byte  `json:"file"`
}

func NewFileRequest(from, to *string, fileName string, file []byte, count int64, size int64) (*File, error) {
	return &File{
		From:     from,
		To:       to,
		FileName: fileName,
		Count:    count,
		Size:     size,
		File:     file,
	}, nil
}

func (f File) GenerateRequest() (*Request, error) {
	body, err := json.Marshal(f)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(FileType, body), nil
}
