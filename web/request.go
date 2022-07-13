package web

import (
	"encoding/json"
	"net/http"
	"time"

	"log-prober-go/common"
	"log-prober-go/handler"
)

func request() {
	configuration := common.Configuration()
	response, err := http.Get(configuration.SaAddr + "/go")
	if err != nil {
		common.Logger.Error(err)
		return
	}
	if response.StatusCode == http.StatusOK {
		// 初始化请求变量结构
		formData := make(map[string]any, 10)
		// 调用json包的解析，解析请求body
		json.NewDecoder(response.Body).Decode(&formData)
		for key, value := range formData {
			switch value.(type) {
			case string:
				handler.Patterns[key] = value.(string)
			}
		}
	}

}

func AutoRequestParseConfig() {
	ticker := time.NewTicker(time.Second * 60)
	defer ticker.Stop()
	for {
		request()
		common.Logger.Debug("AutoRequestParseConfig success ....")
		<-ticker.C
	}

}
