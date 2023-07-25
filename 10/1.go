/*
实现阻塞读且并发安全的map
如何实现map中key不存在get操作等待，直到key存在或者超时，保证并发安全

思路：
1. 挂起goroutine并唤醒，并发安全 用channel (并发安全，实现阻塞)
2. select case 控制超时
3. sync.RWMutex
*/
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type sp interface {
	//存入key /val,如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Wt(key string, val interface{})

	//读取一个key，如果key不存在阻塞，等待key存在或者超时
	Rd(key string, timeout time.Duration) interface{}
}
type mySafeMap struct {
	//每个key都需要锁
	info map[string]*myStruct
	//加锁
	sync.RWMutex
}
type myStruct struct {
	value interface{}
	//通道
	ch    chan bool
	isHad bool
}

//写入key
func (m *mySafeMap) Wt(key string, val interface{}) {
	m.Lock()
	defer m.Unlock()
	//判断key值是否存在
	if v, ok := m.info[key]; ok {
		v.value = val
		if !v.isHad {
			close(v.ch)
		}
		v.isHad = true
		//表示channel里面有值
		//v.ch <- true
		//close(v.ch)
	} else {
		mm := &myStruct{
			value: val,
			ch:    make(chan bool, 1),
			isHad: true,
		}
		m.info[key] = mm
		close(mm.ch)
	}
}

//读取key
func (m *mySafeMap) Rd(key string, timeout time.Duration) interface{} {
	m.Lock()
	if v, ok := m.info[key]; ok {
		if v.isHad == true {
			m.Unlock()
			return v.value
		} else {
			m.Unlock()
			select {
			case <-v.ch:
				return v.value
			case <-time.After(timeout):
				fmt.Printf("我们超时了 key= %s \n", key)
				return nil
			}
		}
	} else {
		mm := &myStruct{
			value: nil,
			ch:    make(chan bool, 1),
			isHad: false,
		}
		m.info[key] = mm
		m.Unlock()
		select {
		//获取值
		case <-mm.ch:
			return mm.value
		case <-time.After(time.Second):
			fmt.Printf("我们超时了%s\n", key)
			return nil
		}
	}
}

func main() {
	mm := &mySafeMap{
		info: make(map[string]*myStruct),
	}
	//写协程
	go func() {
		for i := 0; i < 10; i++ {
			//strconv将int转换成string
			mm.Wt(strconv.Itoa(i), i)
			time.Sleep(time.Second)
		}
	}()
	//读协程
	go func() {
		for i := 9; i >= 0; i-- {
			fmt.Println(mm.Rd(strconv.Itoa(i), time.Second))
		}
	}()
	time.Sleep(time.Second * 11)
}
