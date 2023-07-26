package main

import (
	"encoding/json"
	"fmt"
)

//定义People结构体，只有Name属性
type People struct {
	//Name需要是大写
	Name string `json:"name"`
}

func main() {
	js := `{
        "name":"11"
    }`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("people: %#v\n", p)
}
