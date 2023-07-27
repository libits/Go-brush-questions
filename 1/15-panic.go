//panic发生 但是进程不退出
// panic发生执行三个过程
// 1.defer
// 2.调用堆栈
// 3.exit(2)

//foo2发生错误 不退出将panic如下写
func foo2(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println(err)
			panic(error)
		}
	}
}