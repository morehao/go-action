package main

import (
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
)

const (
	DeployModelPhysical = "physical"
	DeployModelK8S      = "kubernetes"
	DeployModelDocker   = "docker"
)

// DetectEnvironment 仅适配 cgroup v2，自动判断 K8s、Docker、物理机
func DetectEnvironment(ctx *gin.Context) string {
	// 优先：K8s 环境变量（非常可靠）
	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		glog.Infof(ctx, "Detected Kubernetes environment via env var")
		return DeployModelK8S
	}

	// 检查 .dockerenv 文件（仅适用于标准 Docker 环境）
	if _, err := os.Stat("/.dockerenv"); err == nil {
		glog.Infof(ctx, "Detected Docker environment via .dockerenv")
		return DeployModelDocker
	}

	// 读取 cgroup v2 信息
	// data, err := os.ReadFile("/proc/1/cgroup")
	// if err != nil {
	// 	glog.Warnf(ctx, "Failed to read /proc/1/cgroup: %v", err)
	// 	return DeployModelPhysical
	// }

	return DeployModelPhysical
}

func isLikelyContainerID(path string) bool {
	segments := strings.Split(path, "/")
	last := segments[len(segments)-1]
	return len(last) >= 30 && len(last) <= 64 && isHex(last)
}

func isHex(s string) bool {
	for _, c := range s {
		if !((c >= 'a' && c <= 'f') || (c >= '0' && c <= '9')) {
			return false
		}
	}
	return true
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
