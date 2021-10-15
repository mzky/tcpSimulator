package main

import (
	"tcpSimulator/common"
	"tcpSimulator/logs"
	"tcpSimulator/server"
	"tcpSimulator/socket"
)

func main() {
	p := common.Parameter()
	logs.InitLog(*p.LogFile)

	if *p.Long { // 参数 -l 开启socket长连接
		socket.Long(p.Address, *p.Timeout)
	}

	if *p.Socket { //　参数 -s 开启socket短连接
		socket.Short(p.Address)
	}

	if *p.Udp { //　参数 -s 开启socket短连接
		server.Udp(p.Address)
	}

	//　默认执行
	server.Http(p.Address, *p.Interface, *p.CertPath)

}
