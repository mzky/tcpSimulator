package socket

import (
	"fmt"
	"net"
	"tcpSimulator/common"
	"time"

	"github.com/sirupsen/logrus"
)

func Long(host string, timeout int) {
	netListen, err := net.Listen("tcp", host)
	common.CheckError(err)
	defer netListen.Close()
	logrus.Infof("Running long-lived socket server")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		logrus.Infof(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn, timeout)
	}
}

func handleConnection(conn net.Conn, timeout int) { //处理连接
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		common.CheckError(err)
		message := make(chan byte)
		go HeartBeating(conn, message, timeout) //心跳计时
		go GravelChannel(buf, message)          //检测每次Client是否有数据传来
		fmt.Println(string(buf[:n]))
		result := common.TimeNow() + "_" + common.GetLocalIp() + "_result:" + string(buf[:n])
		conn.Write([]byte(result))
	}
}

func HeartBeating(conn net.Conn, message chan byte, timeout int) { //心跳计时，根据GravelChannel判断Client是否在设定时间内发来信息
	select {
	case <-message:
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Minute))
		break
	case <-time.After(time.Duration(timeout) * time.Second):
		logrus.Errorf("超时!!!")
		conn.Close()
	}
}

func GravelChannel(n []byte, mess chan byte) {
	for _, v := range n {
		mess <- v
	}
	close(mess)
}
