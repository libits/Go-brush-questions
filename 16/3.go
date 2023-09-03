/*
生产者/消费者模式
goroutinr+channel实现消费者 生产者模式
storageChan 存储生产的商品管道
shopChan 存储转运给商店的商品管道
exitChan判断主线程结束的管道
*/
package main

import (
	"fmt"
	"strconv"
	"time"
)

type Product struct {
	Name string
}

func main() {
	storageChan := make(chan Product, 10)
	shopChan := make(chan Product, 1)
	exitChan := make(chan bool, 1)

	for i := 1; i < 9; i++ {
		go Producer(storageChan, 10)
	}
	go Logistics(storageChan, shopChan)
	go Consumer(shopChan, 10, exitChan)
	if <-exitChan {
		return
	}

	fmt.Println()
}
func Producer(storagrChan chan Product, count int) {
	for {
		producer := Product{"商品:" + strconv.Itoa(count)}
		storagrChan <- producer
		count--
		time.Sleep(time.Second)
		fmt.Println("生产了", producer)
		if count < 1 {
			return

		}
	}

}
func Logistics(storageChan chan Product, shopChan chan<- Product) {
	for {
		product := <-storageChan
		shopChan <- product
		fmt.Println("运输了", product)
	}
}
func Consumer(shopChan <-chan Product, count int, exitChan chan<- bool) {
	for {
		product := <-shopChan
		fmt.Println("消费了", product)
		count--
		if count < 1 {
			exitChan <- true
			return
		}
	}
}
