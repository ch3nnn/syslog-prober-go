package handler

import (
	"time"

	json "github.com/bytedance/sonic"
	re "github.com/dlclark/regexp2"
	"log-prober-go/domain"
	"log-prober-go/libs/go-syslog/format"
	"log-prober-go/output/kafka"
	"log-prober-go/utils"
)

// Patterns 正则表达式 TODO 内存变量获取解析配置
var Patterns = make(map[string]string)

type RegexpParseHandler struct {
	logParts       format.LogParts
	syslogProtocol *domain.SyslogProtocol
}

func NewRegexpParseHandler(logParts format.LogParts) *RegexpParseHandler {
	return &RegexpParseHandler{logParts: logParts, syslogProtocol: new(domain.SyslogProtocol)}
}

func (h RegexpParseHandler) Handle() (message []byte) {
	marshal, _ := json.Marshal(h.logParts)
	if err := json.Unmarshal(marshal, &h.syslogProtocol); err != nil {
		panic(err)
	}
	var parseMsg = make(map[string]string, 10)
	if len(Patterns) != 0 {
		compile, err := re.Compile(Patterns["config"], 0)
		if err != nil {
			panic(err)
		}
		if m, _ := compile.FindStringMatch(h.syslogProtocol.Src); m != nil {
			for i, g := range m.Groups() {
				if i != 0 {
					parseMsg[g.Name] = g.Captures[0].String()
				}
			}
		}
	}
	message, _ = json.Marshal(
		domain.Message{
			SrcMsg:         h.syslogProtocol.Src,
			ProtocolMsg:    h.syslogProtocol,
			ParseMsg:       parseMsg,
			ParseTimestamp: time.Now().UnixMilli(),
		})
	return message

}

func (h RegexpParseHandler) Output(message []byte) {
	// 在 goroutine 中恢复程序 一定要 recover 异常捕获 记录错误
	defer utils.Recover()
	producer := kafka.Producer()
	producer.Send(message)
}
