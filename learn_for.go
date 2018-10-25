package main

import "fmt"

func main() {
	sum := 1
	// for i := 0;i<10;i++ {
	//   sum += i
	// }
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
	other()
}

// 其中 i:=0 i++可以省略 甚至可以省略前后两个分号
// 也就是说 在for语句可以当做while语句

func other() {

	var b int = 15
	var a int
	//定义数据，不够的话补0
	numbers := [16]int{1, 2, 3, 5}
	fmt.Print(numbers)

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}
