package response

import (
	"encoding/json"
	"errors"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/request"
)

const DownloadFileType = "read-file"

type DownloadFile struct {
	FileName string `json:"file_name"`
	Count    int64  `json:"count"`
	Size     int64  `json:"size"`
	File     []byte `json:"file"`
}

func NewDownloadFileResponse(fileName string, file []byte, count int64, size int64) *DownloadFile {
	return &DownloadFile{
		FileName: fileName,
		Count:    count,
		Size:     size,
		File:     file,
	}
}

func (f DownloadFile) GenerateResponse() (*request.Request, error) {
	body, err := json.Marshal(f)
	if err != nil {
		return nil, errors.New("couldn't marshal body: " + err.Error())
	}

	return New(DownloadFileType, body), nil
}
