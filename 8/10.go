package main

type People interface {
	Speak(string) string
}
type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "111" {
		talk = "222"
	} else {
		talk = "hi"
	}
	return
}
func main() {
	var peo People = Student{}
	//var peo People = &Student{}
	think := "111"
	println(peo.Speak(think))
}

//go tool compile -l -p main 10.go
//go tool nm 10.go |grep T
