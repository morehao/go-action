package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gcontext/gincontext"
	"github.com/morehao/golib/glog"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

// 健康检查
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"service":   "server-fingerprint",
		"version":   "1.0.0",
		"timestamp": time.Now().Unix(),
	})
}

// GetSystemInfoHandler 获取服务器系统信息
func GetSystemInfoHandler(ctx *gin.Context) {
	fingerprint, err := getServerInfo()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get server fingerprint",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   fingerprint,
	})
}

// getServerInfo 获取服务器信息
func getServerInfo() (*SystemInfoResponse, error) {
	fp := &SystemInfoResponse{
		Timestamp: time.Now().Unix(),
	}

	// 基础系统信息
	fp.Hostname, _ = os.Hostname()
	fp.Platform = runtime.GOOS
	fp.OS = runtime.GOOS
	fp.Arch = runtime.GOARCH
	fp.Environment = DetectEnvironment()

	// 获取主机信息
	if hostInfo, err := host.Info(); err == nil {
		fp.HostID = hostInfo.HostID
		fp.KernelVersion = hostInfo.KernelVersion
	}

	// CPU 信息
	if cpuInfo, err := cpu.Info(); err == nil && len(cpuInfo) > 0 {
		fp.CPUInfo = cpuInfo[0].ModelName
	}
	fp.CPUCores = runtime.NumCPU()

	// 内存信息
	if memInfo, err := mem.VirtualMemory(); err == nil {
		fp.Memory = memInfo.Total / 1024 / 1024 / 1024 // GB
	}

	// 网络信息
	fp.MACAddresses = GetMACAddresses()

	// 机器 ID
	fp.MachineID = GetMachineID()

	// 磁盘信息
	fp.DiskInfo = GetDiskInfo()

	// 容器/K8s 特定信息
	switch fp.Environment {
	case "docker":
		fp.ContainerID = GetContainerID()
	case "kubernetes":
		fp.ContainerID = GetContainerID()
		fp.PodName = os.Getenv("HOSTNAME") // 在 K8s 中通常是 Pod 名
		fp.NodeName = os.Getenv("NODE_NAME")
		fp.Namespace = os.Getenv("POD_NAMESPACE")
	}

	// 生成指纹
	fp.Fingerprint = generateFingerprint(fp)

	return fp, nil
}

// 生成指纹
func generateFingerprint(fp *SystemInfoResponse) string {
	// 组合关键信息生成指纹
	key := fmt.Sprintf("%s-%s-%s-%s-%s",
		fp.MachineID,
		fp.HostID,
		fp.CPUInfo,
		strings.Join(fp.MACAddresses, ","),
		fp.Platform,
	)

	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])[:16] // 取前16位作为指纹
}

func GetFingerprint(ctx *gin.Context) {
	var fingerprint string

	deployMode := DetectEnvironment()
	switch deployMode {
	case DeployModelPhysical:
		hostInfo, err := host.Info()
		if err != nil {
			glog.Errorf(ctx, "[GetFingerprint] get hostname error: %v", err)
			gincontext.Fail(ctx, err)
			return
		}
		fingerprint = hostInfo.HostID
	}
	gincontext.Success(ctx, &FingerprintResponse{
		Fingerprint: fingerprint,
	})
}
