package main

import (
	"fmt"
	"time"
)

func main() {
	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}

	//异步读取
	go func() {
		// for a := range abc {
		// 	fmt.Println("a: ", a)
		// }
		for {
			select {
			case a, ok := <-abc:
				if ok {
					fmt.Println("A:", a)
				} else {
					fmt.Println("A", a)
				}
			}
		}
	}()
	close(abc)
	fmt.Println("close")
	time.Sleep(time.Second * 1)
}

/*
channel关闭之后仍会消费值
close
a:  0
a:  1
a:  2
a:  3
a:  4
a:  5
a:  6
a:  7
a:  8
a:  9
*/
