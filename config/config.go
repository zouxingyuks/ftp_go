package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var configLogs *logrus.Entry
var Configs *viper.Viper

func ParseConfig() {
	configLogs = NewLog("configs")
	// 指定配置文件路径
	configDir := "./config/"
	configName := "config"
	configType := "yaml"
	// 如果 Configs 目录不存在，则创建它
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err = os.MkdirAll(configDir, 0755); err != nil {
			configLogs.Errorln(err)
		}
	}
	// 解析配置文件
	Configs.AddConfigPath(configDir)
	Configs.SetConfigName(configName)
	Configs.SetConfigType(configType)
	// 配置文件出错

	if err := Configs.ReadInConfig(); err != nil {
		// 如果找不到配置文件，则提醒生成配置文件并创建它
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			configPath := configDir + configName + "." + configType
			fmt.Printf("[warning] Config file not found. Generating default config file at %s\n", configPath)
			if err := Configs.WriteConfigAs(configPath); err != nil {
				configLogs.Errorf("[error] Failed to generate default config file. %s", err)
			}
			// 再次读取配置文件
			if err := Configs.ReadInConfig(); err != nil {
				configLogs.Errorf("[error] Failed to read config file. %s", err)
			}
		} else {
			// 配置文件被找到，但产生了另外的错误
			configLogs.Errorf("[error] Failed to parse config file. %s", err)
		}
	}
}

func LoadDefaultConfig() {
	defaultConfig := map[string]interface{}{
		"# 下面是默认配置文件": nil,
		"host": map[string]interface{}{
			"ip":   "0.0.0.0",
			"port": "9090",
		},
		"logs": map[string]interface{}{
			"level": "info",
			"path":  "./logs/logrus.log",
		},
		"dir": map[string]interface{}{
			"root": "/home",
			"tmp":  "/tmp",
		},
		"user": map[string]interface{}{
			"root": map[string]interface{}{
				"mod":      "root",
				"password": "root",
			},
			"default": map[string]interface{}{
				"mod":      "default",
				"password": "123456",
			},
		},
	}
	//初始化viper
	Configs = viper.New()
	//将默认值设置到config中
	Configs.MergeConfigMap(defaultConfig)
}
