package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type OSInfo struct {
	OSType   string
	OSArch   string
	Hostname string
}

func (o OSInfo) PrintInfo() {
	fmt.Println("os type ==>", o.OSType)
	fmt.Println("arch ==>", o.OSArch)
	fmt.Println("hostname ==>", o.Hostname)
}

type RAMInfo struct {
	Total       uint64
	Available   uint64
	Used        uint64
	UsedPercent float64
}

func (r RAMInfo) PrintInfo() {
	fmt.Println("total RAM ==>", r.Total)
	fmt.Println("available RAM ==>", r.Available)
	fmt.Println("used RAM ==>", r.Used)
	fmt.Println("RAM usage ==>", r.UsedPercent)
}

type DiskInfo struct {
	Device    string
	TotalSize uint64
	FreeSize  uint64
}

func (d DiskInfo) PrintInfo() {
	fmt.Println("Device ==>", d.Device)
	fmt.Println("Total Size ==>", d.TotalSize, "GB")
	fmt.Println("Free Size ==>", d.FreeSize, "GB")
}

type CPUInfo struct {
	Modelname string
	Cores     int32
}

func (c CPUInfo) PrintInfo() {
	fmt.Println("CPU Model ==>", c.Modelname)
	fmt.Println("Cores ==>", c.Cores)
}

func main() {

	for {
		var input string
		fmt.Println("What information do you want?(os, ram, disk, cpu, all)")
		fmt.Scanln(&input)
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "os":
			// OS INFORMATION
			oSInfo := OSInfo{
				OSType: runtime.GOOS,
				OSArch: runtime.GOARCH,
			}

			hostName, err := os.Hostname()

			if err != nil {
				fmt.Printf("Error getting hostname: %v\n", err)
				return
			}
			oSInfo.Hostname = hostName

			oSInfo.PrintInfo()

		case "ram":
			// RAM INFORMATION
			vmStat, err := mem.VirtualMemory()
			if err != nil {
				fmt.Printf("Error getting memory info: %v\n", err)
				return
			}

			ramInfo := RAMInfo{
				Total:       vmStat.Total / 1024 / 1024,
				Available:   vmStat.Available / 1024 / 1024,
				Used:        vmStat.Used / 1024 / 1024,
				UsedPercent: vmStat.UsedPercent,
			}

			ramInfo.PrintInfo()

		case "disk":
			// DISK INFORMATION
			partitions, err := disk.Partitions(false)
			if err != nil {
				fmt.Printf("Error getting disk partitions: %v\n", err)
				return
			}
			var diskInfos []DiskInfo
			for _, partition := range partitions {
				diskInfo, err := disk.Usage(partition.Mountpoint)
				if err != nil {
					fmt.Printf("Error getting disk usage info for %s: %v\n", partition.Mountpoint, err)
					continue
				}
				diskInfos = append(diskInfos, DiskInfo{
					Device:    partition.Device,
					TotalSize: diskInfo.Total / 1024 / 1024 / 1024,
					FreeSize:  diskInfo.Free / 1024 / 1024 / 1024,
				})
			}
			for _, diskInfo := range diskInfos {
				diskInfo.PrintInfo()
			}

		case "cpu":
			// CPU INFORMATION
			cpuInfos, err := cpu.Info()
			if err != nil {
				fmt.Printf("Error getting CPU info: %v\n", err)
				return
			}
			var cpuInfoList []CPUInfo
			for _, cpuInfo := range cpuInfos {
				cpuInfoList = append(cpuInfoList, CPUInfo{
					Modelname: cpuInfo.ModelName,
					Cores:     cpuInfo.Cores,
				})
			}
			for _, cpuInfo := range cpuInfoList {
				cpuInfo.PrintInfo()
			}
		case "all":
			oSInfo := OSInfo{
				OSType: runtime.GOOS,
				OSArch: runtime.GOARCH,
			}

			hostName, err := os.Hostname()
			if err != nil {
				fmt.Printf("Error getting hostname: %v\n", err)
				return
			}
			oSInfo.Hostname = hostName
			oSInfo.PrintInfo()

			vmStat, err := mem.VirtualMemory()
			if err != nil {
				fmt.Printf("Error getting memory info: %v\n", err)
				return
			}

			ramInfo := RAMInfo{
				Total:       vmStat.Total / 1024 / 1024,
				Available:   vmStat.Available / 1024 / 1024,
				Used:        vmStat.Used / 1024 / 1024,
				UsedPercent: vmStat.UsedPercent,
			}
			ramInfo.PrintInfo()

			partitions, err := disk.Partitions(false)
			if err != nil {
				fmt.Printf("Error getting disk partitions: %v\n", err)
				return
			}
			var diskInfos []DiskInfo
			for _, partition := range partitions {
				diskInfo, err := disk.Usage(partition.Mountpoint)
				if err != nil {
					fmt.Printf("Error getting disk usage info for %s: %v\n", partition.Mountpoint, err)
					continue
				}
				diskInfos = append(diskInfos, DiskInfo{
					Device:    partition.Device,
					TotalSize: diskInfo.Total / 1024 / 1024 / 1024,
					FreeSize:  diskInfo.Free / 1024 / 1024 / 1024,
				})
			}
			for _, diskInfo := range diskInfos {
				diskInfo.PrintInfo()
			}

			cpuInfos, err := cpu.Info()
			if err != nil {
				fmt.Printf("Error getting CPU info: %v\n", err)
				return
			}
			var cpuInfoList []CPUInfo
			for _, cpuInfo := range cpuInfos {
				cpuInfoList = append(cpuInfoList, CPUInfo{
					Modelname: cpuInfo.ModelName,
					Cores:     cpuInfo.Cores,
				})
			}
			for _, cpuInfo := range cpuInfoList {
				cpuInfo.PrintInfo()
			}
		default:
			fmt.Println("Invalid input, please enter one of the options: 'os', 'ram', 'disk', 'cpu', 'all'.")
		}

		var moreInfo string
		fmt.Println("Do you want to display more information? (yes/no):")
		fmt.Scanln(&moreInfo)
		moreInfo = strings.TrimSpace(strings.ToLower(moreInfo))

		if moreInfo != "yes" {
			fmt.Println("goodbye :)")
			break
		}
	}

}
