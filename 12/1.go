/*
要求每秒钟调用一次proc并保证程序不退出(recover)
*/
package main

import (
	"fmt"
	"time"
)

func proc() {
	panic("ok")
}
func main() {
	go func() {
		timer := time.NewTicker(time.Second)
		for {
			select {
			case <-timer.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Printf("我们捕获了异常：%s\n", err)
						}
					}()
					proc()
				}()
			}
		}

	}()
	select {}
}
