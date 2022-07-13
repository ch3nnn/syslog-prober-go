package main

import (
	"log"

	"github.com/bytedance/gopkg/util/gopool"
	"log-prober-go/common"
	"log-prober-go/handler"
	"log-prober-go/libs/go-syslog"
	"log-prober-go/utils"
	"log-prober-go/web"
)

type LogProberApplication struct {
	channel   syslog.LogPartsChannel
	server    *syslog.Server
	parseName int
}

func init() {
	// 加载配置文件
	if err := common.LoadConfiguration("config.yaml"); err != nil {
		log.Println(err)
		return
	}
	common.LogInitialization() // 加载日志
	// common.CacheInitialization()    // 初始化redis池
	go web.AutoRequestParseConfig() // 自动请求解析配置
}

//
// NewLogProberApplication
//  @Description: 日志探针应用程序
//  @param channel 日志频道
//  @param server 日志服务
//  @return *LogProberApplication
//
func NewLogProberApplication(channel syslog.LogPartsChannel, server *syslog.Server) *LogProberApplication {
	return &LogProberApplication{channel: channel, server: server}
}

//
// Run
//  @Description: 启动方法
//  @receiver a LogProberApplication 对象
//  @param addr 本地服务地址端口
//
func (a LogProberApplication) Run(addr string) {
	channelHandler := syslog.NewChannelHandler(a.channel)
	a.server.SetFormat(syslog.Automatic)
	a.server.SetHandler(channelHandler)
	a.server.ListenUDP(addr)
	a.server.Boot()
	common.Logger.Infof("Server %s Running....", addr)
	gopool.Go(func() {
		defer utils.Recover()
		for logParts := range a.channel {
			requestHandler := handler.NewRequestHandler(logParts)
			requestHandler.Handle()
		}
	})
	a.server.Wait()

}

func main() {
	application := NewLogProberApplication(make(syslog.LogPartsChannel), syslog.NewServer())
	application.Run(common.Configuration().ServerAddr)
}
