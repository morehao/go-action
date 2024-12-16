package utils

import (
	"github.com/gin-gonic/gin"
	"net"
)

func GetLocalIp() string {
	addrs, _ := net.InterfaceAddrs()
	var ip string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				if ip != "127.0.0.1" {
					return ip
				}
			}
		}
	}
	return "127.0.0.1"
}

func GetClientIp(ctx *gin.Context) (clientIP string) {
	if ctx == nil {
		return clientIP
	}
	return ctx.ClientIP()
}
