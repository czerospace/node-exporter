package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"net/http"
	"node-exporter-with-consul/global"
)

func InitRegister() {
	// 服务注册
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成对应的检查对象

	check := &api.AgentServiceCheck{
		HTTP:     fmt.Sprintf("http://%s:%d/health", global.ExporterIP, global.LocalPort),
		Timeout:  "5s",
		Interval: "10s",
	}
	fmt.Println(check)
	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = "node-exporter-with-consul"
	// 将服务器 ip 作为 uuid 注册到 consul 中
	registration.ID = global.ExporterIP
	registration.Port = global.LocalPort
	fmt.Println("global.ServerConfig.Tags is : ", global.ServerConfig.Tags)
	registration.Tags = global.ServerConfig.Tags
	registration.Address = global.ExporterIP
	registration.Check = check

	// 注册
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	// 增加一个注销 consul 接口
	http.HandleFunc("/deregister", func(w http.ResponseWriter, r *http.Request) {
		if err := client.Agent().ServiceDeregister(global.ExporterIP); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("无法注销服务: %s", err)))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("服务已注销"))
	})
}
