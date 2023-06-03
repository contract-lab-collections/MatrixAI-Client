package cpu

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
)

// 定义 CPUInfo 结构体
type CPUInfo struct {
	ModelName string  // CPU 名称
	Cores     int32   // CPU 核心数
	Mhz       float64 // CPU 频率，单位 MHz
}

// 获取 CPU 信息并返回 CPUInfo 结构体切片
func GetCPUInfo() ([]CPUInfo, error) {
	cpuInfos := make([]CPUInfo, 0)
	cpuInfoStats, err := cpu.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU info: %w", err)
	}

	for _, info := range cpuInfoStats {
		cpuInfos = append(cpuInfos, CPUInfo{
			ModelName: info.ModelName,
			Cores:     info.Cores,
			Mhz:       info.Mhz,
		})
	}
	return cpuInfos, nil
}
