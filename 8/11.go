// nil指针结构体返回给interface类型
//代码输出是什么

package main

import (
	"fmt"
)

//非空接口
type Peoplex interface {
	Show()
}

type Studentx struct{}

func (stu *Studentx) Show() {

}

//func live() *Studentx{
func live() Peoplex {
	//只是声明没有赋值
	var stu *Studentx
	return stu
}

func main() {
	tmp := live()
	if tmp == nil {
		fmt.Println("AAAAAAA", tmp)
	} else {
		fmt.Println("BBBBBBB", tmp)
	}
}

/*
AAAAAAA <nil>
BBBBBBB <nil>
*/
