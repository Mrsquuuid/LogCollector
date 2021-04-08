package main

import (
	"LogCollector/D4/log_transfer/es"
	"LogCollector/D4/log_transfer/kafka"
	"LogCollector/D4/log_transfer/model"
	"fmt"
	"github.com/go-ini/ini"
)

// log transfer
// 从kafka消费日志数据,写入ES

func main(){
	// 1. 加载配置文件
	//得到config结构体对应的一个指针。
	var cfg = new(model.Config)
	//拿到指针后，在加载的时候直接传cfg即可。
	err := ini.MapTo(cfg, "./config/logtransfer.ini")
	if err != nil {
		fmt.Printf("load config failed,err:%v\n", err)
		panic(err)
	}
	fmt.Printf("%#v\n", *cfg)
	fmt.Println("load config success")
	// 2. 连接ES
	// 2.1 初始化一个es连接的client
	// 2.2 对外提供一个往es里写入的函数
	err = es.Init(cfg.ESConf.Address, cfg.ESConf.Index, cfg.ESConf.GoNum, cfg.ESConf.MaxSize)
	if err != nil {
		fmt.Printf("Init es failed,err:%v\n", err)
		panic(err)
	}
	fmt.Println("Init ES success")
	// 3. 连接kafka
	// 3.1 连接kafka，创建分区消费者。
	// 3.2 每个分区消费者分别取出数据，通过sendtoes发给es。
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		fmt.Printf("connect to kafka failed,err:%v\n", err)
		panic(err)
	}
	fmt.Println("Init kafka success")
	// 在这儿停顿!
	select {}
}
