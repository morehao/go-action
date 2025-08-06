package main

// SystemInfoResponse 服务器指纹信息结构
type SystemInfoResponse struct {
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

type FingerprintResponse struct {
	Fingerprint string `json:"fingerprint"` // 指纹信息
}

type K8SClusterInfo struct {
	UID string `json:"uid"`
}
