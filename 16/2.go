/*
goroutine和Channel实现所有2000之内的素数
intChan所有的数字通道
primeChan 素数的通道
exitChan  true
*/
package main

import (
	"fmt"
)

var intChan chan int = make(chan int, 100)

func main() {
	var primeChan chan int = make(chan int, 100)
	var exitChan chan bool = make(chan bool, 8)
	go initChan(20000)
	for i := 0; i <= 8; i++ {
		go isPrimeA(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i <= 8; i++ {
			<-exitChan
		}
		close(primeChan)
	}()
	for {
		v, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("素数：", v)
	}
}
func initChan(num int) {
	for i := 1; i < num; i++ {
		intChan <- i

	}
	close(intChan)
}
func isPrimeA(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		flag = true
		if !ok {
			break
		}
		for j := 2; j < num; j++ {
			if num%j == 0 {
				flag = false
				continue
			}

		}
		if flag {
			primeChan <- num
		}
	}
	exitChan <- true
}
