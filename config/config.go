package config

// ServerConfig 配置文件结构体
type ServerConfig struct {
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
	NetInfo    NetConfig    `mapstructure:"linux" json:"linux"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NetConfig struct {
	Interface string `mapstructure:"interface" json:"interface"`
	Subnet    string `mapstructure:"subnet" json:"subnet"`
}
