// 由于子切片和母切片共享底层内存 下述这样写会导致母切片的底层内存无法被释放
func subSlice() []byte {
	longSlice := [1 << 20]byte{}
	shortSlice := longSlice[3:7]
	return shortSlice
}

//正确写法
func subSlice1() []byte {
	longSlice := [1 << 20]byte{}
	var shortSlice []int = make([]int, 10)
	shortSlice := longSlice[3:7]
	return shortSlice
}