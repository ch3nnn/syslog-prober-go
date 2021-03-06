package kafka

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"log-prober-go/common"
)

const (
	// 超时
	kafkaTimeOut = time.Second * 5
	// Sync Async kafka生产者发送信息方式
	Sync  = "Sync"
	Async = "Async"
)

var (
	addressError           = errors.New("kafka address is error")          // 地址错误
	newSyncProducerError   = errors.New("sarama.NewSyncProducer err")      // 创建同步生产者错误
	newAsyncProducerError  = errors.New("sarama.NewAsyncProducer err")     // 创建异步生产者错误
	asyncProducerNullError = errors.New("sarama.NewSyncProducer is null")  // 异步生产者为空
	sendMessageError       = errors.New("send kafka message err")          // 发送信息错误
	modeError              = errors.New("kafka mode error  sync or async") // 模式错误

)

var producer AbsKafkaProducer

// KafkaConfig kafka生产者
type KafkaConfig struct {
	addressList []string       // 地址列表
	topic       string         // kafka topic
	config      *sarama.Config // kafka配置信息
}

// NewProducer 新建kafka生产者并选择同步方式
// 一般使用异步方式
func newProducer(address []string, topic string, duration time.Duration, syncOrAsync string) AbsKafkaProducer {
	var newProducer AbsKafkaProducer
	switch syncOrAsync {
	case Sync:
		newProducer = &SyncKafkaProducer{}
	case Async:
		newProducer = &AsyncKafkaProducer{}
	default:
		panic(modeError)
	}
	newProducer.NewKafkaProducer(address, topic, duration)
	return newProducer
}

// NewProducerByMessage 创建kafka基本生产者
func NewProducerByMessage(address []string, topic string, duration time.Duration) *KafkaConfig {

	// 根据字符串解析地址列表
	if len(address) < 1 || address[0] == "" {
		panic(addressError)
	}
	// 配置producer参数
	sendConfig := sarama.NewConfig()
	sendConfig.Producer.Return.Successes = true
	sendConfig.Producer.Timeout = kafkaTimeOut
	sendConfig.Producer.RequiredAcks = sarama.NoResponse
	if duration != 0 {
		sendConfig.Producer.Timeout = duration
	}
	return &KafkaConfig{
		addressList: address,
		topic:       topic,
		config:      sendConfig,
	}
}

// Producer 获取生产者
func Producer() AbsKafkaProducer {
	if producer == nil {
		producer = newProducer(
			common.Configuration().KafkaServers,
			common.Configuration().KafkaTopic,
			time.Second*5,
			common.Configuration().KafkaProMode,
		)
	}
	return producer
}
