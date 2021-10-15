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
	reader := bufio.NewReader(os.Stdin)
	line, err1 := reader.ReadString('\n')
	common.CheckError(err1)
	line = strings.TrimSpace(line)
	if line == "exit" {
		fmt.Println("客户端退出..")
		return
	}
	_, err = conn.Write([]byte(line + "\n"))
	common.CheckError(err)
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	common.CheckError(err2)
	fmt.Println(string(buf[:n]))
}
