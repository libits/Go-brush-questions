package main

import "fmt"

func main() {
	//make(num.len,cap)
	s := make([]int, 0, 3)
	prevCap := cap(s)
	for i := 0; i < 100; i++ {
		s = append(s, 1)
		currcap := cap(s)
		if currcap > prevCap {
			fmt.Println(prevCap, currcap)
			prevCap = currcap
		}
	}

}

/*
每次扩容都是二倍
3 6
6 12
12 24
24 48
48 96
96 192
*/
