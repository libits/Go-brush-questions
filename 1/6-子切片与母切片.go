package main

import "fmt"

func main() {
	//0 0 0 nil nil
	arr := make([]int, 3, 5)

	//0 0 len=2  cap=4
	brr := arr[1:3]
	brr[0] = 4
	brr = append(brr, 4)
	fmt.Println(arr, brr)
}

/*
[0 4 0] [4 0 4]
*/
