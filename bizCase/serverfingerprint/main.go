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
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

// Fingerprint 服务器指纹信息结构
type Fingerprint struct {
	// 基础信息
	Hostname    string `json:"hostname"`
	Platform    string `json:"platform"`
	Environment string `json:"environment"`

	// 系统信息
	OS            string `json:"os"`
	Arch          string `json:"arch"`
	KernelVersion string `json:"kernel_version"`

	// 硬件信息
	CPUInfo  string `json:"cpu_info"`
	CPUCores int    `json:"cpu_cores"`
	Memory   uint64 `json:"memory_gb"`

	// 网络信息
	MACAddresses []string `json:"mac_addresses"`

	// 唯一标识
	MachineID string `json:"machine_id"`
	HostID    string `json:"host_id"`

	// 磁盘信息
	DiskInfo []DiskInfo `json:"disk_info"`

	// 容器/K8s 信息
	ContainerID string `json:"container_id,omitempty"`
	PodName     string `json:"pod_name,omitempty"`
	NodeName    string `json:"node_name,omitempty"`
	Namespace   string `json:"namespace,omitempty"`

	// 生成的指纹
	Fingerprint string `json:"fingerprint"`
	Timestamp   int64  `json:"timestamp"`
}

type DiskInfo struct {
	Device string `json:"device"`
	Total  uint64 `json:"total_gb"`
	Free   uint64 `json:"free_gb"`
}

// 生成指纹
func generateFingerprint(fp *Fingerprint) string {
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

// 获取服务器指纹信息
func getServerFingerprint() (*Fingerprint, error) {
	fp := &Fingerprint{
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

// API 处理函数
func getFingerprintHandler(c *gin.Context) {
	fingerprint, err := getServerFingerprint()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get server fingerprint",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   fingerprint,
	})
}

// 健康检查
func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"service":   "server-fingerprint",
		"version":   "1.0.0",
		"timestamp": time.Now().Unix(),
	})
}

func main() {
	// 设置 Gin 模式
	gin.SetMode(gin.ReleaseMode)

	// 创建 Gin 引擎
	r := gin.Default()

	// 添加中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 路由设置
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server Fingerprint Service",
			"endpoints": []string{
				"GET /health - 健康检查",
				"GET /fingerprint - 获取服务器指纹",
			},
		})
	})

	r.GET("/health", healthHandler)
	r.GET("/fingerprint", getFingerprintHandler)

	// 获取端口配置
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 启动服务
	fmt.Printf("🚀 Server starting on port %s\n", port)
	fmt.Printf("📝 Environment: %s\n", DetectEnvironment())

	if err := r.Run(":" + port); err != nil {
		fmt.Printf("❌ Failed to start server: %v\n", err)
		os.Exit(1)
	}
}
