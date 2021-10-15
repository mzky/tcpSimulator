package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"tcpSimulator/common"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	common.CheckError(err)
	fmt.Println("connect success")
	send(conn)
}

func send(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		common.CheckError(err)
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出..")
			break
		}
		_, err = conn.Write([]byte(line + "\n"))
		common.CheckError(err)

		buf := make([]byte, 1024)
		n, err1 := conn.Read(buf)
		common.CheckError(err1)
		fmt.Println(string(buf[:n]))
	}
}
