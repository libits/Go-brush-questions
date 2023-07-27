package main

import "fmt"

func foo() (i int) {
	i = 9
	defer func() {
		fmt.Println(i) //5
	}()
	defer fmt.Println(i) //9
	return 5

}
func main() {
	foo()
}

/*
9
5
*/
