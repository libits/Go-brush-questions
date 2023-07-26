package main

import "fmt"

const (
	a = iota //0
	b = iota //1
)
const (
	name = "menglu"
	c    = iota //1
	age  = 100  //1
	d    = iota //3
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

/*
0
1
1
3
*/
