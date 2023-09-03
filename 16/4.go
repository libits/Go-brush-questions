/*
定时执行任务，类似延时消息队列
周期性执行某个任务
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	//方式1
	// fmt.Println("当前时间：", time.Now())
	// timer := time.NewTimer(time.Second * 3)
	// t := <-timer.C  //timer.C是一个只读的管道
	// fmt.Println(t)

	//方式2
	t := time.After(time.Second * 3)
	fmt.Println(t)
}
