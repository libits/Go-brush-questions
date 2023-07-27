//出错原因：有可能for执行完 协程还未执行 无法打印
for i := 0; i < 10; i++ {
	go func() {
		fmt.Println(i)
	}()

}
//解决办法：传参
for i := 0; i < 10; i++ {
	go func() {
		fmt.Println(i)
	}(i)

}