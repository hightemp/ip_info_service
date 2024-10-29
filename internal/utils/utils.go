package utils

import (
	"encoding/binary"
	"net"
)

func IpStringToInt(ip string) uint32 {
	return Ip2int(net.ParseIP(ip))
}

func Ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func Int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
