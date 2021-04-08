package kafka

import (
	"LogCollector/D4/log_transfer/es"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
)

// 初始化kafka连接
// 从kafka里面取出日志数据


func Init(addr []string, topic string)(err error){
	// golang如何使用sarama访问kafka
	// 1.创建新的消费者
	consumer, err:= sarama.NewConsumer(addr, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	// 2.拿到指定topic下面的所有分区列表
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList{ // 3. 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		var pc sarama.PartitionConsumer
		//传入topic、分区号、最新的偏移量。
		pc, err = consumer.ConsumePartition(topic,  int32(partition),sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		//defer pc.AsyncClose()
		// 4. 异步从每个分区消费信息
		fmt.Println("start to consume...")
		go func(sarama.PartitionConsumer){
			fmt.Println("in sarama.PartitionConsumer")
			//这里也可以直接读一条往es里写一条，相当于把连接kafka的功能和连接es功能通过函数调用连接到一起了
			//这样就会变成一个同步过程。我们为了提高系统的健壮性，使用channel。
			//把消息取出来放到channel里，另外一头从channel取，变成了一个异步关系。
			for msg:=range pc.Messages(){
				//logDataChan<-msg // 为了将同步流程异步化,所以将取出的日志数据先放到channel中
				//channel待会儿写到es那边（两别皆可）
				fmt.Println(msg.Topic, string(msg.Value))
				var m1 map[string]interface{}
				err = json.Unmarshal(msg.Value, &m1)
				if err != nil {
					fmt.Printf("unmarshal msg failed, err:%v\n", err)
					continue
				}
				//es包里的方法
				es.PutLogData(m1)
			}
		}(pc)
	}
	return
}