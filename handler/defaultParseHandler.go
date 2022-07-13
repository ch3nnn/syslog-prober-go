package handler

import (
	"log-prober-go/libs/go-syslog/format"
)

type DefaultParseHandler struct {
	logParts format.LogParts
}

func (d DefaultParseHandler) Handle() (message []byte) {
	// TODO implement me
	panic("implement me")
}

func (d DefaultParseHandler) Output(message []byte) {
	// TODO implement me
	panic("implement me")
}

func NewDefaultParseHandler(logParts format.LogParts) *DefaultParseHandler {
	return &DefaultParseHandler{logParts: logParts}
}
