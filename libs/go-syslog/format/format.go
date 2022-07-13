package format

import (
	"bufio"
	"time"

	"log-prober-go/libs/go-syslog/internal/syslogparser"
)

type LogParts map[string]any

type LogParser interface {
	Parse() error
	Dump() LogParts
	Location(*time.Location)
}

type Format interface {
	GetParser([]byte) LogParser
	GetSplitFunc() bufio.SplitFunc
}

type parserWrapper struct {
	syslogparser.LogParser
}

func (w *parserWrapper) Dump() LogParts {
	return LogParts(w.LogParser.Dump())
}
