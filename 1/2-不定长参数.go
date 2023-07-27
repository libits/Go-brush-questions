package main

import "fmt"

//args为slice int类型的不定长参数...
func foo1(s string, args ...int) {
	//len(args)
	//args[0]
	var arr []int
	arr = args
	fmt.Println(arr)
}
func main() {
	arr := []int{}
	foo1("A", arr...)
	//不定长参数用append居多
	s := append([]rune("ABC"))
	//s := append([]byte("ABC"))
}
