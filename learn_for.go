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
}

// 其中 i:=0 i++可以省略 甚至可以省略前后两个分号
// 也就是说 在for语句可以当做while语句

// 死循环写法

func test() {
	for {

	}
}
