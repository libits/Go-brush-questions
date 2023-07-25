/*高并发下的锁与map的读写
在一个高并发的web服务器中，要限制ip的频繁访问，模拟100个ip同时并发访问服务器，每个ip要重复访问1k次
每个ip三分钟之内只能访问一次，要求成功输出success:100
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban1 struct {
	visitIPs map[string]time.Time
	sync.RWMutex
}

func NewBan1() *Ban1 {
	return &Ban1{visitIPs: make(map[string]time.Time)}
}
func (o *Ban1) visit(ip string) bool {
	o.Lock()
	defer o.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	//没有值 则开新协程 并且开启定时器
	o.visitIPs[ip] = time.Now()
	go o.delTimeOut(ip)
	return false
}

func (o *Ban1) delTimeOut(ip string) {
	ticker := time.NewTicker(time.Second * 180)
	select {
	//只读
	case <-ticker.C:
		o.Lock()
		defer o.Unlock()
		fmt.Printf("我执行了删除key：%s \n", ip)
		delete(o.visitIPs, ip)
	}
}
func main() {
	var success int64 = 0
	Ban1 := NewBan1()
	wait := sync.WaitGroup{}
	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()
				// randd := rand.Int63n(2)
				// time.Sleep(time.Millisecond * time.Duration(randd))
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !Ban1.visit(ip) {
					// fmt.Printf("查看randd的值：%d \n", randd)
					//原子操作：解决多线程表示的并发问题，防止出现success：99
					//success传的是变量执行的地址
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}

	}
	wait.Wait()
	fmt.Println("success:", success)
}
