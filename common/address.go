package common

import (
	"github.com/mzky/tls"
	"net"
	"strings"
)

func GetLocalIP() string {
	ipList, _ := GetLocalIPList()
	var addrList []string
	for _, s := range ipList {
		if s != "127.0.0.1" && tls.IsIPv4(s) {
			addrList = append(addrList, s)
		}
	}
	return strings.Join(addrList, ",")
}

func appendIPNet(slice []net.IPNet, element net.IPNet) []net.IPNet {
	if element.IP.IsLinkLocalUnicast() { // ignore link local IPv6 address like "fe80::x"
		return slice
	}

	return append(slice, element)
}

func GetLocalIpNets() (map[string][]net.IPNet, error) {
	iFaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	returnMap := make(map[string][]net.IPNet)
	for _, iFace := range iFaces {
		if iFace.Flags&net.FlagUp == 0 { // Ignore down adapter
			continue
		}

		address, err := iFace.Addrs()
		if err != nil {
			continue
		}

		ipNets := make([]net.IPNet, 0)
		for _, addr := range address {
			switch v := addr.(type) {
			case *net.IPAddr:
				ipNets = appendIPNet(ipNets, net.IPNet{v.IP, v.IP.DefaultMask()})
			case *net.IPNet:
				ipNets = appendIPNet(ipNets, *v)
			}
		}
		returnMap[iFace.Name] = ipNets
	}

	return returnMap, nil
}

func GetLocalIPList() ([]string, error) {
	ipArray := make([]string, 0)
	ipMap, err := GetLocalIpNets()
	if err != nil {
		return nil, err
	}
	mapAddr := make(map[string]string) //去重
	for _, ipNets := range ipMap {
		for _, ipNet := range ipNets {
			mapAddr[ipNet.IP.String()] = ipNet.IP.String()
		}
	}

	for _, ip := range mapAddr {
		ipArray = append(ipArray, strings.TrimSpace(ip))
	}
	return ipArray, nil
}
