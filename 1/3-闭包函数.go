package main

import (
	"fmt"
)

func sub() func() {
	//i闭包，内部变量i与外部变量相同
	i := 10
	b := func() {
		i--
		fmt.Printf("%n", i)

	}
	return b
}
func main() {
	sub()
}
