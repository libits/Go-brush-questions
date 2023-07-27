package main

import "fmt"

var i interface{}

func main() {
	i = 3
	switch i.(type) {
	case int:
		fmt.Println(i)
	case byte:
		fmt.Println(i)
	}
}
