package main

import (
	"MatrixAI-Client/hardwareinfo"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli"
)

var Version = "v0.0.1"

func setupApp() *cli.App {
	app := cli.NewApp()
	app.Usage = "MatrixAI-Client"
	app.Action = startService
	app.Version = Version
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{}
	app.Before = func(context *cli.Context) error {
		fmt.Println("NumCPU:", runtime.NumCPU())

		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func main() {
	if err := setupApp().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startService(context *cli.Context) error {

	hwInfo, err := hardwareinfo.GetHardwareInfo()
	if err != nil {
		log.Fatalf("Error getting hardware info: %v", err)
	}

	fmt.Printf("Hardware Info:\n%+v\n", hwInfo)

	// //调用cpu内的GetCPUInfo方法获取CPU信息
	// cpuInfos, err := cpu.GetCPUInfo()
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return err
	// }

	// for idx, cpuInfo := range cpuInfos {
	// 	fmt.Printf("------ CPU #%d ------\n", idx+1)
	// 	fmt.Printf("CPU Model Name: %s\n", cpuInfo.ModelName)
	// 	fmt.Printf("CPU Cores Number: %d\n", cpuInfo.Cores)
	// 	fmt.Printf("CPU Frequency: %.2f MHz\n", cpuInfo.Mhz)
	// }

	// //调用disk内的GetDiskInfo方法获取disk信息
	// diskInfos, err := disk.GetDiskInfo()
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return err
	// }

	// for idx, diskInfo := range diskInfos {
	// 	fmt.Printf("------ disk #%d ------\n", idx+1)
	// 	fmt.Printf("disk Path: %s\n", diskInfo.Path)
	// 	fmt.Printf("disk TotalSpace: %.2f GB\n", diskInfo.TotalSpace)
	// 	fmt.Printf("disk FreeSpace: %.2f GB\n", diskInfo.FreeSpace)
	// }

	// //调用memory内的GetMemoryInfo方法获取memory信息
	// memoryInfo, err := memory.GetMemoryInfo()
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return err
	// }

	// fmt.Printf("------ memory ------\n")
	// fmt.Printf("memory Total: %.2f GB\n", memoryInfo.TotalMemory)

	// //调用gpu内的GetIntelGPUInfo方法获取gpu信息
	// gpuInfos, err := gpu.GetIntelGPUInfo()
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return err
	// }

	// for idx, gpuInfo := range gpuInfos {
	// 	fmt.Printf("------ gpu #%d ------\n", idx+1)
	// 	fmt.Printf("gpu Model Name: %s\n", gpuInfo.Model)
	// }

	return nil
}
