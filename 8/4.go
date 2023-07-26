package main

type People1 struct{}

func (p *People1) ShowA() {
	println("showA")
	// fmt.Printf("当前对象：%#v \n", p)
	p.ShowB()
}
func (p *People1) ShowB() {
	println("showB")
}

type Teacher struct {
	People1
	//People2
}

func (t *Teacher) ShowB() {
	println("teacher showB")
}

type People2 struct{}

func (p *People2) ShowA() {
	println("show People2 A")
	// fmt.Printf("当前对象：%#v \n", p)
	p.ShowB()
}
func (p *People2) ShowB() {
	println("show People2 B")
}

func main() {
	t := Teacher{}

	//println("获取teacher中People地址", &t.People2)

	t.ShowA()
}
