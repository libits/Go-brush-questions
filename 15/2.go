package main

import (
	"fmt"
	"sync"
	"time"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func(), wait *sync.WaitGroup) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

func main() {
	wait := sync.WaitGroup{}
	for j := 0; j < 10000; j++ {
		wait.Add(10000)
		d := new(Once)

		f := func() {
			fmt.Printf("我执行了")

		}
		for i := 0; i < 10000; i++ {
			go d.Do(f, &wait)
		}
		wait.Wait()
		fmt.Printf("结束： \n")
		if j%50 == 0 {
			time.Sleep(time.Second)
		}

	}
	time.Sleep(time.Millisecond * 1000)

}
