/*
用goroutine和channel
1. 开启一个writeData协程，向管道intChan写入50个整数
2. 开启一个readData协程，从管道intChan中读取writeData写入的数据
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var intChan chan int

func writeData(intChan chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		var tempInt int
		tempInt = rand.Intn(4) + 10
		fmt.Printf("写入%v,第%v次写\n", tempInt, i)
		intChan <- tempInt
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {
	var count int
	for {
		v, ok := <-intChan
		count++
		if !ok {
			break
		}
		fmt.Printf("%v,%v\n", v, count)
	}
	exitChan <- true
	close(exitChan)
}
func main() {
	intChan = make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	time.Sleep(time.Second * 2)
	fmt.Println("end")
}
