package initialize

import (
	"github.com/spf13/viper"

	"node-exporter-with-consul/global"
)

// InitConfig 初始化配置文件
func InitConfig() {
	// new 得到一个 viper
	v := viper.New()
	// 设置配置文件路径
	configFileName := "config/config.yaml"
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 实例化一个 ServerConfig{} 结构体
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
}
