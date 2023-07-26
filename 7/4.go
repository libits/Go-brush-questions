package main

import "fmt"

type People struct {
	Name string
}

//*People结构体 ；String()方法，返回string类型  String()自己调用自己 导致栈溢出
func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People{}
	fmt.Println(p.String())
}
