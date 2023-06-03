package memory

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

// 定义 MemoryInfo 结构体
type MemoryInfo struct {
	TotalMemory float64 // 总内存大小，单位 GB
}

// 获取内存信息并返回 MemoryInfo 结构体
func GetMemoryInfo() (MemoryInfo, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return MemoryInfo{}, fmt.Errorf("failed to get memory info: %w", err)
	}

	return MemoryInfo{
		TotalMemory: float64(vmStat.Total) / (1 << 30),
	}, nil
}
