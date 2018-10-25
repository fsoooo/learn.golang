package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("result:", add_1(15, 25))
	fmt.Println("result:", add_2(15, 25))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println("杨辉三角形:")
	ShowYangHuiTriangle()

	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数作为值 */
	fmt.Println(getSquareRoot(17))

}

/**
Go 语言函数定义格式如下：
func function_name( [parameter list] ) [return_types] {
   函数体
}
 */

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

//行数
const LINES int = 10

func ShowYangHuiTriangle() {
	nums := []int{}
	for i := 0; i < LINES; i++ {
		//补空白
		for j := 0; j < (LINES - i); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < (i + 1); j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)
			fmt.Print(value, " ")
		}
		fmt.Println("")
	}
}
