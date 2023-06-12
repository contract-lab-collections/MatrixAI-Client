package hardwareinfo

import (
	"MatrixAI-Client/hardwareinfo/cpu"
	"MatrixAI-Client/hardwareinfo/disk"
	"MatrixAI-Client/hardwareinfo/flops"
	"MatrixAI-Client/hardwareinfo/gpu"
	"MatrixAI-Client/hardwareinfo/machine_uuid"
	"MatrixAI-Client/hardwareinfo/memory"
)

// HardwareInfo 结构体用于存储所有硬件(cpu\disk\flops\gpu\memory)信息
type HardwareInfo struct {
	MachineUUID machine_uuid.MachineUUID `json:"MachineUUID"` // 机器 UUID
	CPUInfos    []cpu.InfoCPU            `json:"CPUInfos"`    // CPU 信息
	DiskInfos   []disk.InfoDisk          `json:"DiskInfos"`   // 硬盘信息
	MemoryInfo  memory.InfoMemory        `json:"InfoMemory"`  // 内存信息
	GPUInfos    []gpu.InfoGPU            `json:"GPUInfos"`    // GPU 信息（仅限英特尔显卡）
	FlopsInfo   flops.InfoFlop           `json:"InfoFlop"`    // FLOPS 信息
}

// GetHardwareInfo 函数收集并返回全部硬件信息
func GetHardwareInfo() (HardwareInfo, error) {
	var hwInfo HardwareInfo

	// 获取 UUID 信息
	machineUUID, err := machine_uuid.GetInfoMachineUUID()
	if err != nil {
		return hwInfo, err
	}
	hwInfo.MachineUUID = machineUUID

	// 获取 CPU 信息
	cpuInfo, err := cpu.GetCPUInfo()
	if err != nil {
		return hwInfo, err
	}
	hwInfo.CPUInfos = cpuInfo

	// 获取硬盘信息
	diskInfos, err := disk.GetDiskInfo()
	if err != nil {
		return hwInfo, err
	}
	hwInfo.DiskInfos = diskInfos

	// 获取内存信息
	memInfo, err := memory.GetMemoryInfo()
	if err != nil {
		return hwInfo, err
	}
	hwInfo.MemoryInfo = memInfo

	// 获取 GPU 信息（仅限英特尔显卡）
	gpuInfos, err := gpu.GetIntelGPUInfo()
	if err != nil {
		return hwInfo, err
	}
	hwInfo.GPUInfos = gpuInfos

	// 获取 FLOPS 信息
	if len(cpuInfo) > 0 {
		numCores := int(cpuInfo[0].Cores)
		flopsInfo := flops.GetFlopsInfo(numCores)
		hwInfo.FlopsInfo = flopsInfo
	}

	return hwInfo, nil
}
