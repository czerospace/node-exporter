package global

import (
	"node-exporter-with-consul/config"
)

var (
	// ServerConfig 定义全局 config 变量
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	ExporterIP   string
)
