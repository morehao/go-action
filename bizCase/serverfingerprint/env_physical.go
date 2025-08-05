package main

import (
	"bytes"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/shirou/gopsutil/v4/disk"
)

// Linux: 读取 /etc/machine-id 或 /var/lib/dbus/machine-id
func getMachineIDLinux() string {
	paths := []string{
		"/etc/machine-id",
		"/var/lib/dbus/machine-id",
	}

	for _, path := range paths {
		if data, err := os.ReadFile(path); err == nil {
			id := strings.TrimSpace(string(data))
			if id != "" {
				return id
			}
		}
	}
	return "unknown"
}

// macOS: 使用 IOPlatformUUID
func getMachineIDDarwin() string {
	out, err := exec.Command("ioreg", "-rd1", "-c", "IOPlatformExpertDevice").Output()
	if err != nil {
		return "unknown"
	}
	re := regexp.MustCompile(`"IOPlatformUUID" = "([^"]+)"`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) == 2 {
		return matches[1]
	}
	return "unknown"
}

// Windows: 使用 wmic 命令获取 UUID（主板唯一标识）
func getMachineIDWindows() string {
	cmd := exec.Command("wmic", "csproduct", "get", "UUID")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "unknown"
	}
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "UUID") || line == "" {
			continue
		}
		return line
	}
	return "unknown"
}

// 获取 MAC 地址
func GetMACAddresses() []string {
	var macs []string
	interfaces, err := net.Interfaces()
	if err != nil {
		return macs
	}

	for _, iface := range interfaces {
		if iface.HardwareAddr != nil && iface.HardwareAddr.String() != "" {
			// 排除本地回环和虚拟接口
			if !strings.HasPrefix(iface.Name, "lo") &&
				!strings.HasPrefix(iface.Name, "docker") &&
				!strings.HasPrefix(iface.Name, "veth") {
				macs = append(macs, iface.HardwareAddr.String())
			}
		}
	}
	return macs
}

// 获取磁盘信息
func GetDiskInfo() []DiskInfo {
	var disks []DiskInfo
	partitions, err := disk.Partitions(false)
	if err != nil {
		return disks
	}

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}

		disks = append(disks, DiskInfo{
			Device: partition.Device,
			Total:  usage.Total / 1024 / 1024 / 1024, // GB
			Free:   usage.Free / 1024 / 1024 / 1024,  // GB
		})
	}
	return disks
}
