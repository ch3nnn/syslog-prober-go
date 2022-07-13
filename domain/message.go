package domain

//
// Message
//  @Description: kafka 写入字段
//
type Message struct {
	SrcMsg          string            `json:"srcMsg"`          // 源 syslog 字符串日志
	ProtocolMsg     *SyslogProtocol   `json:"protocolMsg"`     // syslog 标准字段
	ParseMsg        map[string]string `json:"parseMsg"`        // 解析字段
	ParseTemplateId int64             `json:"parseTemplateId"` // 解析模板 id
	ProbeId         int64             `json:"probeId"`         // 探针器 id
	ParseTimestamp  int64             `json:"parseTimestamp"`  // 解析时间戳
	AssetGroup      string            `json:"AssetGroup"`      // 资产组名称

}
