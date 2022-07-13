package common

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"log-prober-go/domain"
)

var configuration *domain.Configuration

//
// LoadConfiguration
//  @Description: 加载配置文件
//  @param path
//  @return error
//
func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config domain.Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	configuration = &config
	return err
}

//
// Configuration
//  @Description: 获取配置选项
//  @return *Configuration
//
func Configuration() *domain.Configuration {
	return configuration
}
