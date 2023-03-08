# tcpSimulator
提供模拟TCP/UDP服务端 支持socket长短连接和http服务

其中udp服务返回syslog信息
```
# ./tcpSimulator_linux_amd64 -h
Usage: ./tcpSimulator_linux_amd64 [OPTIONS] args

OPTIONS
  -a string
        ip address (default "0.0.0.0")
  -c string
        Https TLS certificate path
  -f string
        Log file (default "./tcpdemo.log")
  -i string
        Start http server and set path (default "/*any")
  -l    Start long-lived socket server
  -p string
        Port (default "9000")
  -s    Start short socket server
  -t int
        Set overtime time (default 30)
  -u    Start Udp server

USAGES
  Custom Address
        ./tcpSimulator_linux_amd64 -p 9000
        ./tcpSimulator_linux_amd64 -u -a 192.168.1.100 -p 9000
  Short Socket Server
        ./tcpSimulator_linux_amd64 -s
        ./tcpSimulator_linux_amd64 -s -p 9000
  Long-Lived Socket Server
        ./tcpSimulator_linux_amd64 -s -l
        ./tcpSimulator_linux_amd64 -s -l -p 9000
        ./tcpSimulator_linux_amd64 -s -l -p 9000 -t 30
  UDP Server(SYSLOG)
        ./tcpSimulator_linux_amd64 -u
        ./tcpSimulator_linux_amd64 -u -p 9000
  HTTP Server
        ./tcpSimulator_linux_amd64
        ./tcpSimulator_linux_amd64 -i
        ./tcpSimulator_linux_amd64 -i -p 9000
        ./tcpSimulator_linux_amd64 -i /custom_path
  HTTPS Server
        .. Default TLS certificate file name: server.pem  server.key
        ./tcpSimulator_linux_amd64 -c .
        ./tcpSimulator_linux_amd64 -c /certificate_path/
        ./tcpSimulator_linux_amd64 -c . -p 9000
        ./tcpSimulator_linux_amd64 -c . -p 9000 -i /custom_path

```
