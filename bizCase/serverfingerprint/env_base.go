package main

import (
	"os"
	"runtime"
	"strings"
)

const (
	DeployModelPhysical = "physical"
	DeployModelK8S      = "kubernetes"
	DeployModelDocker   = "docker"
)

// 检测运行环境
func DetectEnvironment() string {
	// 检查是否在 Kubernetes 中
	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		return "kubernetes"
	}

	// 检查是否在 Docker 中
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return "docker"
	}

	// 检查 cgroup 信息
	if data, err := os.ReadFile("/proc/1/cgroup"); err == nil {
		content := string(data)
		if strings.Contains(content, "docker") {
			return "docker"
		}
		if strings.Contains(content, "kubepods") {
			return "kubernetes"
		}
	}

	return "physical"
}

// GetMachineID 返回当前机器的唯一标识（多平台适配）
func GetMachineID() string {
	switch runtime.GOOS {
	case "linux":
		return getMachineIDLinux()
	case "darwin":
		return getMachineIDDarwin()
	case "windows":
		return getMachineIDWindows()
	default:
		return "unknown"
	}
}
