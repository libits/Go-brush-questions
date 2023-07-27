package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	for i, ele := range arr {
		arr[i] += 1
		ele += 8
		fmt.Println(arr[i], ele)
	}
}

/*
2 9
3 10
4 11
*/
