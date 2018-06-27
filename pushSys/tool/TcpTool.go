package tool

import (
	"net"
	"fmt"
)

func GetIp() (string, error) {
	adds, err := net.InterfaceAddrs()
	if err != nil {
		return "", fmt.Errorf("cannot read ip")
	}
	for _, address := range adds {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.To4().String(), nil
			}
		}
	}
	return "", fmt.Errorf("no ip")
}
