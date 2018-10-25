package main

import (
	"fmt"
)

func main() {
	fmt.Println("result:", add_1(15, 25))
	fmt.Println("result:", add_2(15, 25))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

//函数：关键字/函数名/参数列表/返回值类型 ，ps:参数类型写在参数值后面
func add_1(x int, y int) int {
	return x + y
}

// 当多个连续的参数类型相同时，可以只在最后写一个即可
func add_2(x, y int) int {
	return x + y
}

// 返回多个值
func swap(x, y string) (string, string) {
	return y, x
}
