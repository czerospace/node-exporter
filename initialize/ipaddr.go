package initialize

import (
	"fmt"
	"net"
	"node-exporter-with-consul/global"
)

func InitGetIpaddr() {
	iface, err := net.InterfaceByName(global.ServerConfig.NetInfo.Interface)
	if err != nil {
		panic(err)
	}

	addrs, err := iface.Addrs()
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {

		switch v := addr.(type) {
		case *net.IPNet:
			if v.IP.To4() != nil && isBondSubnet(v.IP) {
				global.ExporterIP = v.IP.String()
				fmt.Printf("addr is ################## %s\n", addr)
				fmt.Println("IPNET global.ExporterIP is" + global.ExporterIP)
			}
		}
	}
}

func isBondSubnet(ip net.IP) bool {
	// 请根据您的特定需求修改此处子网掩码
	_, subnet, err := net.ParseCIDR(global.ServerConfig.NetInfo.Subnet)
	if err != nil {
		panic(err)
	}

	return subnet.Contains(ip)
}
