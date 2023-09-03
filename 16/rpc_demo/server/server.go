package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type MathService struct {
}

func (m *MathService) Add(arg [2]int, reply *int) error {
	*reply = arg[0] + arg[1]
	return nil
}
func main() {
	//创建MathServices实例
	mathService := new(MathService)

	//注册服务
	err := rpc.Register(mathService)
	if err != nil {
		//
		fmt.Println(err)
	}
	//监听
	listener, err := net.Listen("tcp", "localhost:1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)

	}
}
