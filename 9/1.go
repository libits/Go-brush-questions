// golang协程和channel配合使用
// 实现两个goroutine,其中一个产生随机数并写入go channel中，
//  另外一个从channel中读取数字并打印到标准输出，最后输出五个随机数
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int)
	wait := sync.WaitGroup{}

	wait.Add(5)
	go func() {
		defer close(ch)
		for i := 1; i < 6; i++ {
			ch <- rand.Intn(100)
		}
	}()

	go func() {
		for {
			select {
			case chInt, ok := <-ch:
				if ok {
					fmt.Printf("打印channel值： %d \n", chInt)
					wait.Done()
				}
			}
		}
	}()
	wait.Wait()

}
