package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	//stu是一个变量，地址固定不变错误写法
	/*for _,stu :=range stu{
		m.[stu.name]=&stu
	}*/
	for _, stu := range stus {
		tmpStu := stu
		m[stu.Name] = &tmpStu
	}

	fmt.Printf("我们m的值是： %#v \n", m)
	fmt.Printf("我们m的值是： %#v \n", m["wang"])
	fmt.Printf("我们m的值是： %#v \n", m["li"])
	fmt.Printf("我们m的值是： %#v \n", m["zhou"])
}

func main() {
	pase_student()
}
