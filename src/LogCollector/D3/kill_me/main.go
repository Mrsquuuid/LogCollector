package main

import "sync"

// 1 1 2 3
func fib(n int64) int64{
	if n <= 2{
		return 1
	}
	return fib(n-1) + fib(n-2)
}


func main(){
	var wg sync.WaitGroup
	wg.Add(100)
	for i:=0;i<100;i++{
		go fib(10000)
	}
	wg.Wait()
}
