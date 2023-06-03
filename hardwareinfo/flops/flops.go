package flops

import (
	"math/rand"
	"sync"
	"time"
)

const numOperations = 10000000 // 定义每个执行单元的浮点运算次数

// FlopsInfo 结构体用于存储 FLOPS 计算结果
type FlopsInfo struct {
	Flops float64 // 每秒浮点运算次数
}

// floatOperation 执行一定数量的浮点数运算
func floatOperation() {
	a := rand.Float64()
	b := rand.Float64()

	for i := 0; i < numOperations; i++ {
		_ = a * b
	}
}

// GetFlopsInfo 函数计算 FLOPS 并返回包含结果的 FlopsInfo 结构体，接受 CPU 核心数作为入参
func GetFlopsInfo(cpuCores int) FlopsInfo {
	var wg sync.WaitGroup
	wg.Add(cpuCores)

	startTime := time.Now()

	for i := 0; i < cpuCores; i++ {
		go func() {
			floatOperation()
			wg.Done()
		}()
	}

	wg.Wait()

	duration := time.Since(startTime)
	totalOperations := numOperations * cpuCores
	flops := float64(totalOperations) / duration.Seconds()

	return FlopsInfo{Flops: flops}
}
