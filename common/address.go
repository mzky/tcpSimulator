package common

import "net"

func GetLocalIp() string {
	address, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range address {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}

	return ""

}
