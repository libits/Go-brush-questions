package main

import "fmt"

func main() {
	var s string
	test(s)
	var ss interface{}
	test(ss)
}

type student struct {
	Name string
}

/*
case不能指针类型和值类型在一起进行判断 ，必须分开
*/
func test(v interface{}) {
	switch msg := v.(type) {
	case *student:
		_ = msg.Name
		fmt.Println("是*student类型\n")
	case student:
		_ = msg.Name
		fmt.Println("是student类型\n")
	default:
		fmt.Printf("类型是：%T\n", msg)
	}
}
