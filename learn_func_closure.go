package main

import "fmt"

func main() {
	fmt.Println("学习闭包函数")
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	add_func := add(1,2)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())
}

/**
定义一个闭包函数
 */
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/**
 闭包使用方法
 */
func add(x1, x2 int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, x1 + x2
	}
}
