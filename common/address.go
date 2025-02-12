package common

import (
	"github.com/mzky/tls"
	"github.com/safchain/ethtool"
	"net"
	"sort"
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

func queryLinkState(adapterName string) bool {
	et, err := ethtool.NewEthtool()
	if err != nil {
		return false
	}
	defer et.Close()

	if state, _ := et.LinkState(adapterName); state != 1 {
		return false
	}
	info, _ := et.DriverInfo(adapterName)
	if info.NStats == 0 {
		return false
	}

	return true
}

func GetLocalIpNets() (map[string][]net.IPNet, error) {
	iFaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ipNetArray := make(map[string][]net.IPNet)
	for _, iFace := range iFaces {
		if iFace.Flags&net.FlagUp == 0 { // Ignore down adapter
			continue
		}
		if iFace.Flags&net.FlagLoopback == net.FlagLoopback { // Ignore loop back adapter
			continue
		}
		if iFace.HardwareAddr == nil {
			continue
		}
		if queryLinkState(iFace.Name) {
			continue
		}
		kIgnoreAdapterPrefixes := []string{"lo", "tun", "vir"}
		// if adapter name start with lo vir or tun(defined by kIgnoreAdapterPrefixes), we ignore it
		for _, ignoreName := range kIgnoreAdapterPrefixes {
			if strings.HasPrefix(iFace.Name, ignoreName) {
				continue
			}
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
		ipNetArray[iFace.Name] = ipNets
	}

	return ipNetArray, nil
}

func GetLocalIPList() ([]string, error) {
	ipArray := make([]string, 0)
	ipMap, err := GetLocalIpNets()
	if err != nil {
		return nil, err
	}
	for _, ipNets := range ipMap {
		for _, ipNet := range ipNets {
			ipArray = append(ipArray, strings.TrimSpace(ipNet.IP.String()))
		}
	}

	sort.Sort(sort.StringSlice(ipArray))
	return ipArray, nil
}
