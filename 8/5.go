// select case选择是否有优先级
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	//通道内赋值
	int_chan <- 1
	string_chan <- "hello"
	//读取，没有优先级
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
