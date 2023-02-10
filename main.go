package main

import (
	"os"
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
		os.Exit(0)
	}

	if *p.Socket { //　参数 -s 开启socket短连接
		socket.Short(p.Address)
		os.Exit(0)
	}

	if *p.Udp { //　参数 -s 开启socket短连接
		server.Udp(p.Address)
		os.Exit(0)
	}

	//　默认执行
	server.Http(p.Address, *p.Interface, *p.CertPath)
	os.Exit(0)
}
