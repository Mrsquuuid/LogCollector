package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

// cpu info 采集cpu的信息
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	//可能不止一个cpu，用for循环
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率 每秒统计下cpu使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}


// cpu 负载(windows下还没实现)
func getLoad(){
	info, err := load.Avg()
	if err != nil {
		fmt.Printf("load.Avg() failed, err:%v", err)
		return
	}
	fmt.Println(info)
}

// 内存信息
func getMemInfo(){
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed, err:%v", err)
		return
	}
	fmt.Println(info)
}

// host info
func getHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

// 磁盘 信息
func getDiskInfo(){
	// 获取所有分区信息
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get disk partitions failed, err:%v", err)
		return
	}
	fmt.Println(parts)
	//获取每个盘的分区信息
	for _, part := range parts{
		partInfo, err := disk.Usage(part.Mountpoint)
		if err != nil {
			fmt.Printf("get part stat failed, err:%v", err)
			return
		}
		fmt.Println(partInfo)
	}
	// 磁盘IO
	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}


// 网络信息
func getNetInfo(){
	netIOs, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("get net io counters failed, err:%v", err)
		return
	}
	for _, netIO := range netIOs{
		fmt.Println(netIO)
	}
}

func main(){
	//getCpuInfo()
	//getLoad()
	//getMemInfo()
	//getHostInfo()
	getDiskInfo()
	//getNetInfo()
}


