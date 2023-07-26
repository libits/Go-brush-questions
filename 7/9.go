package main

import "fmt"

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"zhouJieLun"}}
	/*
		错误写法map需要hash算法才能获取key，value(位置不固定)
		m["people"].name="li"
	*/
	n := m["people"]
	n.name = "li"
	m["people"] = n
	fmt.Println(n)
	fmt.Println(m)
}
