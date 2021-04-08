package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main(){

	filename := `./xx.log`
	config := tail.Config{
		ReOpen: true,	//日志过大会自动切割
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 2}, //打开这个文件从哪个地方开始读数据
		MustExist: false,  //允许日志不存在
		Poll: true,		//轮询方式
	}

	// 打开文件开始读取数据
	tails, err :=  tail.TailFile(filename, config)
	if err != nil {
		fmt.Println("tail %s failed, err:%v\n", filename, err)
		return
	}

	// 开始读取数据
	var (
		msg *tail.Line
		ok bool
	)
	for {
		msg, ok = <-tails.Lines // chan tail.Line
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n",
				tails.Filename)
			time.Sleep(time.Second) // 读取出错等一秒
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}