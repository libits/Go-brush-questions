package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//连接服务端
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {

	}
	defer client.Close()

	args := [2]int{3, 5}
	var result int
	//调用rpc方法
	err = client.Call("MathService.Add", args, &result)
	if err != nil {

	}
	fmt.Println(result)
}
