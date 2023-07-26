package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	//主协程关闭了 协程无法运行 相当于往关闭的通道写数据
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	//查询数据
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	//defer close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 2)
}
