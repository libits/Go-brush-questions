package main

import "fmt"

type query func(string) string

//...表示slice  数据类型是query
func exec(name string, vs ...query) string {
	ch := make(chan string)
	//匿名函数
	fn := func(i int) {
		ch <- vs[i](name)
	}
	//开启协程
	for i, _ := range vs {
		go fn(i)
	}
	//一旦有值 就关闭
	return <-ch
}

func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}
