package domain

//
// SyslogProtocol
//  @Description: SyslogProtocol 标准字段
//
type SyslogProtocol struct {
	ProbeTime int64  `json:"probeTime"` // 时间戳 (毫秒)
	Hostname  string `json:"hostname"`  // 设备主机名
	Client    string `json:"client"`    // 设备ip
	Tag       string `json:"tag"`       // 标签
	Content   string `json:"content"`   // 内容
	Priority  int64  `json:"priority"`  // 优先权
	Facility  int64  `json:"facility"`  // 设施
	Severity  int64  `json:"severity"`  // 严重性等级
	Src       string `json:"src"`       // 源日志
}
