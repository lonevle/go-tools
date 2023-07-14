package gotools

// 网关相关

import "net"

// HasLocalIP 检测 IP 地址是否是内网地址
func HasLocalIP(addr string) bool {
	ip := net.ParseIP(addr)
	if ip == nil {
		return false
	}

	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4() // (#`O′) nil也可以执行, 伪调用
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}
