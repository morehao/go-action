package main

import (
	"os"
	"strings"
)

// 获取容器 ID
func GetContainerID() string {
	if data, err := os.ReadFile("/proc/self/cgroup"); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if strings.Contains(line, "docker") {
				parts := strings.Split(line, "/")
				if len(parts) > 0 {
					containerID := parts[len(parts)-1]
					if len(containerID) >= 12 {
						return containerID[:12] // 返回短 ID
					}
				}
			}
		}
	}
	return ""
}
