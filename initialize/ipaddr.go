package initialize

import (
	"net"
	"node-exporter-with-consul/global"
)

func InitGetIpaddr() {
	iface, err := net.InterfaceByName(global.ServerConfig.Interfacename)
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
			}
		case *net.IPAddr:
			if v.IP.To4() != nil && isBondSubnet(v.IP) {
				global.ExporterIP = v.IP.String()
			}
		}
	}
}

func isBondSubnet(ip net.IP) bool {
	_, subnet, err := net.ParseCIDR(global.ServerConfig.Subnet) // 请根据您的特定需求修改此处子网掩码
	if err != nil {
		panic(err)
	}

	return subnet.Contains(ip)
}
