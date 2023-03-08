package common

import (
	"flag"
	"fmt"
	"os"
)

type Param struct {
	Port      *string
	Socket    *bool
	Long      *bool
	Interface *string
	Timeout   *int
	Address   string
	IP        *string
	LogFile   *string
	CertPath  *string
	Udp       *bool
	RFC       *string
}

func Parameter() Param {
	var p Param
	p.IP = flag.String("a", "0.0.0.0", "ip address")
	p.Port = flag.String("p", "9000", "Port")
	p.Socket = flag.Bool("s", false, "Start short socket server")
	p.Long = flag.Bool("l", false, "Start long-lived socket server")
	p.Udp = flag.Bool("u", false, "Start Udp server")
	p.Interface = flag.String("i", "/*any", "Start http server and set path")
	p.CertPath = flag.String("c", "", "Https TLS certificate path")
	p.Timeout = flag.Int("t", 30, "Set overtime time")
	p.LogFile = flag.String("f", "./tcpdemo.log", "Log file")

	//p.RFC = flag.String("r", "Automatic", "RFC3164, RFC6587 or RFC5424 or Automatic")

	flag.Usage = usage
	flag.Parse()

	p.Address = fmt.Sprintf("%s:%s", *p.IP, *p.Port)
	return p
}

func doPrintf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, format+"\n", a...)
}

func usage() {
	doPrintf("Usage: %s [OPTIONS] args\n", os.Args[0])
	doPrintf("OPTIONS")
	flag.PrintDefaults()
	doPrintf("")
	doPrintf("USAGES")
	doPrintf("  Custom Address")
	doPrintf("        %s -p 9000", os.Args[0])
	doPrintf("        %s -u -a 192.168.1.100 -p 9000", os.Args[0])
	doPrintf("  Short Socket Server")
	doPrintf("        %s -s", os.Args[0])
	doPrintf("        %s -s -p 9000", os.Args[0])
	doPrintf("  Long-Lived Socket Server")
	doPrintf("        %s -s -l", os.Args[0])
	doPrintf("        %s -s -l -p 9000", os.Args[0])
	doPrintf("        %s -s -l -p 9000 -t 30", os.Args[0])
	doPrintf("  UDP Server(SYSLOG)")
	doPrintf("        %s -u", os.Args[0])
	doPrintf("        %s -u -p 9000", os.Args[0])
	//doPrintf("        %s -u -p 9000 -r RFC5424", os.Args[0])
	//doPrintf("        %s -u -p 9000 -r RFC6587", os.Args[0])
	//doPrintf("        %s -u -p 9000 -r RFC3164", os.Args[0])
	doPrintf("  HTTP Server")
	doPrintf("        %s", os.Args[0])
	doPrintf("        %s -i", os.Args[0])
	doPrintf("        %s -i -p 9000", os.Args[0])
	doPrintf("        %s -i /custom_path", os.Args[0])
	doPrintf("  HTTPS Server")
	doPrintf("        .. Default TLS certificate file name: server.pem  server.key")
	doPrintf("        %s -c .", os.Args[0])
	doPrintf("        %s -c /certificate_path/", os.Args[0])
	doPrintf("        %s -c . -p 9000", os.Args[0])
	doPrintf("        %s -c . -p 9000 -i /custom_path", os.Args[0])
	doPrintf("")
}
