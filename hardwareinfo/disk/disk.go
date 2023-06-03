package disk

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
)

// 定义 DiskInfo 结构体
type DiskInfo struct {
	Path        string  // 磁盘路径（如：/dev/sda1）
	TotalSpace  float64 // 总共的硬盘空间大小，单位 GB
	FreeSpace   float64 // 可用的硬盘空间大小，单位 GB
}

// 获取硬盘信息并返回 DiskInfo 结构体切片
func GetDiskInfo() ([]DiskInfo, error) {
	diskInfos := make([]DiskInfo, 0)
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, fmt.Errorf("failed to get disk info: %w", err)
	}

	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}
		diskInfos = append(diskInfos, DiskInfo{
			Path:       p.Device,
			TotalSpace: float64(usage.Total) / (1 << 30),
			FreeSpace:  float64(usage.Free) / (1 << 30),
		})
	}
	return diskInfos, nil
}