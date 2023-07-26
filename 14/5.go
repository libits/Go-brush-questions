package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func main() {
	a := &Student{Name: "menglu"}
	b := &Student{Name: "menglu"}

	fmt.Println(a == b) //比较指针
	fmt.Printf("a = %s , b = %s \n", a, b)
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})
}
