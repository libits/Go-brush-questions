package main

func main() {
	
}
type Addable interface{
	type int,byte,string
}
func add_test[T Addable](a,b T)T{
	return a+b
}

//go run -gcflags=-G=3 xx.go