package gpu

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// 定义 GPUInfo 结构体
type GPUInfo struct {
	Model string // GPU显卡型号
}

// 获取 Intel GPU 信息并返回一个包含 GPUInfo 结构体的切片
func GetIntelGPUInfo() ([]GPUInfo, error) {
	cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name")
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to get GPU info: %w", err)
	}

	stdoutStr := strings.TrimSpace(stdoutBuf.String())
	gpuLines := strings.Split(stdoutStr, "\n")
	gpuInfos := make([]GPUInfo, 0)

	for _, line := range gpuLines[1:] {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Intel") {
			gpuInfos = append(gpuInfos, GPUInfo{Model: line})
		}
	}

	return gpuInfos, nil
}
