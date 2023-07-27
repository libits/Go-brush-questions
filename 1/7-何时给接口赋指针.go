package main
type Animai interface{
	say()
}
type People struct{

} 

//值 指针
func(c *People)say(){
}
var a Animal
//创建People实例
P:=People{}
//a=P
a=&P