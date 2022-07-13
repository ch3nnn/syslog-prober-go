package handler

import (
	"log-prober-go/domain"
	"log-prober-go/libs/go-syslog/format"
)

type RequestHandler struct {
	parseName int64
	logParts  format.LogParts
}

func NewRequestHandler(logParts format.LogParts) *RequestHandler {
	return &RequestHandler{logParts: logParts}
}

//
// Handle
//  @Description: 根据平台下发规则进行解析正则、json、csv
//  @receiver r RequestHandler
//
func (r RequestHandler) Handle() {
	// TODO 现阶段只有 syslog 采集, 只有正则解析
	/*
		解析模板样例数据
		[
		    {
		        "mode": "REGEX",
		        "field": "message",
		        "regex": "^(?<time>[a-zA-Z]+\\s+[0-9]+\\s+[0-9]+\\:[0-9]+\\:[0-9]+\\s+)(?<device_name>[a-zA-Z]+\\-[a-zA-Z]+\\s+)(?<host>[^\\:\\n]+)\\S*\\s+(?<message>[^\\.\\n]+)\\S*\\s+(?<json>.*)$",
		        "is_timestamp": true
		    },
		    {
		        "mode": "TIME",
		        "field": "TAGno.time",
		        "is_timestamp": true
		    }
		]

	*/

	r.parseName = 3
	var requestHandler ParseHandler
	switch r.parseName {
	case domain.REGEX:
		requestHandler = NewRegexpParseHandler(r.logParts)
	case domain.JSON:
		requestHandler = NewJsonParseHandler(r.logParts)
	default:
		requestHandler = NewDefaultParseHandler(r.logParts)
	}
	// 处理解析
	message := requestHandler.Handle()
	// 解析输出
	go requestHandler.Output(message)

}
