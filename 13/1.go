/*
为sync.WaitGroup中Wait函数支持WaitTimout功能
如果timeout到了超时时间返回true
WaitGroup自然结束返回false
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)

		}(i, c)
		//wg.Wait()
	}
	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}
}
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan bool)
	timer := time.NewTicker(timeout)

	go func() {
		select {
		case <-timer.C:
			ch <- true
		}
	}()
	go func() {
		wg.Wait()
		ch <- false
	}()
	return <-ch
}
