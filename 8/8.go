package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type UserAges struct {
	//map是并发读写不安全的
	ages map[string]int
	//sync.Mutex 互斥锁，只开一个协程
	sync.RWMutex //可并发读 写的时候不能读
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	// ua.Lock()
	// defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {

	u := new(UserAges)
	u.ages = make(map[string]int)

	go mapAdd(u)

	go mapGet(u)

	time.Sleep(time.Second * 1)
}

func mapAdd(userAge *UserAges) {
	for i := 0; i < 100; i++ {
		go userAge.Add(strconv.Itoa(i), i)
	}
}

func mapGet(userAge *UserAges) {
	for i := 0; i < 100; i++ {
		fmt.Println(userAge.Get(strconv.Itoa(i)))
	}
}
