package handler

import (
	"log-prober-go/libs/go-syslog/format"
)

type JsonParseHandler struct {
	logParts format.LogParts
}

func (j JsonParseHandler) Handle() (message []byte) {
	// TODO implement me
	panic("implement me")
}

func (j JsonParseHandler) Output(message []byte) {
	// TODO implement me
	panic("implement me")
}

func NewJsonParseHandler(logParts format.LogParts) *JsonParseHandler {
	return &JsonParseHandler{logParts: logParts}
}
