package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)

	}()
	wg.Wait()
}

/*
出现panic
解释：waitGroup等待一组goroutine完成
主goroutine调用Add方法设置要等待的Goroutine数量
然后每个Goroutine运行并在完成后调用Done方法
同时使用Wait方法阻塞，直到所有Goroutine完成
*/
