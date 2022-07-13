package domain

import "go.uber.org/zap/zapcore"

//
// Configuration
//  @Description: 对应 config.yaml 文件字段
//
type Configuration struct {
	ServerAddr       string        `yaml:"server_addr"`         // 服务地址
	LogLevel         zapcore.Level `yaml:"log_level"`           // 日志等级
	LogPath          string        `yaml:"log_path"`            // 日志路径
	LogMaxSize       int           `yaml:"log_max_size"`        // 文件内容大小, MB
	LogMaxBackups    int           `yaml:"log_max_backups"`     // 保留旧文件最大个数
	LogMaxAge        int           `yaml:"log_max_age"`         // 保留旧文件最大天数
	KafkaServers     []string      `yaml:"kafka_servers"`       // kafka 地址
	KafkaTopic       string        `yaml:"kafka_topic"`         // kafka topic
	KafkaProMode     string        `yaml:"kafka_pro_mode"`      // kafka 生产者异步同步模式
	CacheAddress     string        `yaml:"cache_address"`       // redis 地址
	CacheNetwork     string        `yaml:"cache_network"`       // redis 连接方式
	CachePassword    string        `yaml:"cache_password"`      // redis 密码
	CacheDatabase    int           `yaml:"cache_database"`      // redis 数据库
	SaAddr           string        `yaml:"sa_addr"`             // 平台服务地址
	SaRegularTime    int           `yaml:"sa_regular_time"`     // 定期请求获取配置时间 单位 秒
	SaParseCacheName string        `yaml:"sa_parse_cache_name"` // 解析缓存名称 业务模块名:业务逻辑含义:其他
}
