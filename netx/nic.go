package netx

import (
	"fmt"
	"net"
)

// GetIPv4ByNICName 获取指定网卡的ipv4地址
func GetIPv4ByNICName(name string) (string, error) {
	ift, err := net.InterfaceByName(name)
	if err != nil {
		return "", err
	}

	addrs, err := ift.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.To4() != nil {
			return ipAddr.IP.To4().String(), nil
		}
	}
	return "", fmt.Errorf("no ipv4 address available on '%s'", name)
}

// GetIPv6ByNICName 获取指定网卡的ipv6地址
func GetIPv6ByNICName(name string) (string, error) {
	ift, err := net.InterfaceByName(name)
	if err != nil {
		return "", err
	}

	addrs, err := ift.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.To4() == nil && ipAddr.IP.To16() != nil {
			return ipAddr.IP.To16().String(), nil
		}
	}
	return "", fmt.Errorf("no ipv6 address available on '%s'", name)
}

// GetIPByNICName 获取指定网卡的ip地址（ipv4优先）
func GetIPByNICName(name string) (string, error) {
	ift, err := net.InterfaceByName(name)
	if err != nil {
		return "", err
	}

	addrs, err := ift.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.To4() != nil {
			return ipAddr.IP.To4().String(), nil
		}
		if ipAddr.IP.To16() != nil {
			return ipAddr.IP.To16().String(), nil
		}
	}
	return "", fmt.Errorf("no ipv4/ipv6 address available on '%s'", name)
}
