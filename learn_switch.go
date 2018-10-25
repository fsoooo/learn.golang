package main

import "fmt"

//switch语句
func main() {
	switch a := 10; a {
	case 10:
		fmt.Println("golang switch so bad")
	case 20:
		fmt.Println("hahaha...")
	}
	// 等价于如下写法
	// a := 10
	// switch  {
	// case 10:
	//
	// case:20
	//
	// }
	grade()
}
func grade() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n");
	}
	fmt.Printf("你的等级是 %s\n", grade);
}
