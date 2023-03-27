package config

// ServerConfig 配置文件结构体
type ServerConfig struct {
	ConsulInfo    ConsulConfig `mapstructure:"consul" json:"consul"`
	InterfaceName string       `mapstructure:"interfacename" json:"interfacename"`
	Subnet        string       `mapstructure:"subnet" json:"subnet"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
