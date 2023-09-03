/*
：
如果矩阵中的许多系数都为零， 那么该矩阵就是稀疏的。 对稀疏现象有兴趣是因为它的开发
可以带来巨大的计算节省， 并且在许多大的实践中都会出现矩阵稀疏的问题。
给定一个矩阵， 现在需要逐行和逐列地扫描矩阵， 如果某一行或者某一列内， 存在连续出现
的 0 的个数超过了行宽或者列宽的一半 ， 则认为该行或者该列是稀疏的。
扫描给定的矩阵， 输出稀疏的行数和列数。

*/
package main

import "fmt"

func main() {
	row := 0
	column := 0
	fmt.Scan(&row, &column)
	m := Init(row, column)
	row, column = Cal(m)
	fmt.Println(row)
	fmt.Println(column)
}
func Init(row, column int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < row; i++ {
		r := make([]int, column)
		for j := 0; j < column; j++ {
			fmt.Scan(&r[j])
		}
		result = append(result, r)
	}
	return result
}
func Cal(m [][]int) (int, int) {
	var rCount, cCount int
	var rrefZero, crefZero int
	row := len(m)
	column := len(m[0])
	for i := 0; i < row; i++ {
		rrefZero = 0
		for j := 0; j < column; j++ {
			if m[i][j] == 0 {
				rrefZero += 1
			}
		}
		if rrefZero >= column/2 {
			rCount += 1
		}
	}
	for i := 0; i < column; i++ {
		crefZero = 0
		for j := 0; j < row; j++ {
			if m[j][i] == 0 {
				crefZero += 1
			}
		}
		if crefZero >= row/2 {
			cCount += 1
		}
	}
	return rCount, cCount
}
