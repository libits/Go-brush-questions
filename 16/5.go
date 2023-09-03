//循环时钟
package main

import (
	"fmt"
	"time"
)

func main() {
	var count int = 0
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for {
			//t是一个只读的管道 并将时间取出来
			t := <-ticker.C
			fmt.Println("时间：", t)
			count++
			if count > 2 {
				ticker.Stop()
			} else {
				return
			}
		}
	}()
	time.Sleep(time.Second * 10)

}
