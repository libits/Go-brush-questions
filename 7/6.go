package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	value = 100
	SetValue(102)
	fmt.Println(value)
}

var value int32

func SetValue(delta int32) {
	v := value
	// fmt.Printf("打印： %#v", atomic.CompareAndSwapInt32(&value, v+1, (v+delta)))
	/*for {
		v := value
		// fmt.Printf("打印： %#v", atomic.CompareAndSwapInt32(&value, v, (v+delta)))
		//将&value的值与v进行对比
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
			break
		}
	}
	*/
	atomic.CompareAndSwapInt32(&value, v, (v + delta))
}
