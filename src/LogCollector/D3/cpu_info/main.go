package main

import (
	"fmt"
	"time"
	"github.com/shirou/gopsutil/cpu"
)

// 获取cpu信息


// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", pecent)
	}
}

func main(){

}