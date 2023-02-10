package socket

import (
	"fmt"
	"net"
	"tcpSimulator/common"

	"github.com/sirupsen/logrus"
)

var tcpClientList = make([]net.Conn, 0)

func Short(host string) {
	logrus.Infof("Running short socket server")
	listener, err := net.Listen("tcp", host)
	common.CheckError(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		common.CheckError(err)
		go handleConn(conn)
		tcpClientList = append(tcpClientList, conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			//fmt.Println(err.Error())
			index := -1
			for key, value := range tcpClientList {
				// logrus.Infof("key:%d value:%d\n", key, value)
				if value == conn {
					index = key
				}
			}
			if index != -1 {
				tcpClientList = append(tcpClientList[:index], tcpClientList[index+1:]...)
			}
			return
		}
		fmt.Println(string(buf[:n]))
		result := common.TimeNow() + " # " + common.GetLocalIP() + " # result: " + string(buf[:n])
		conn.Write([]byte(result))
	}

}
