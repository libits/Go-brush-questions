package main

import (
	"fmt"
)

func main() {
	//切片是浅拷贝 ，会改变值 slice是引用传递
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)
}

/*
[a b new]
[a b new]
*/
