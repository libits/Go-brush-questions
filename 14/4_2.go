package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2}

	fmt.Printf("arr1 = %v, &arr1 = %p \n", arr1, &arr1)
	changeArr(arr1)
	fmt.Printf("后——arr1 = %v, &arr1 = %p \n", arr1, &arr1)

	var slice1 = []int{1, 2}
	fmt.Printf("slice = %v, &slice = %p \n", slice1, &slice1)
	changeSlice(slice1)
	fmt.Printf("后——slice = %v, &slice = %p \n", slice1, &slice1)
}

func changeArr(arr [3]int) {
	arr[0] = 1000
	arr[1] = 2000
	fmt.Printf("arr1 = %v, &arr1 = %p \n", arr, &arr)
}

func changeSlice(arr []int) {
	arr[0] = 1000
	arr[1] = 2000
	fmt.Printf("slice = %v, &slice = %p \n", arr, &arr)
}

/*
arr1 = [1 2 0], &arr1 = 0xc000018150
arr1 = [1000 2000 0], &arr1 = 0xc000018180
后——arr1 = [1 2 0], &arr1 = 0xc000018150
slice = [1 2], &slice = 0xc000008078
slice = [1000 2000], &slice = 0xc0000080a8
后——slice = [1000 2000], &slice = 0xc000008078
*/
