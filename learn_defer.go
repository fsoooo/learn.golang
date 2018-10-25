package main

import "fmt"

func main()  {
  defer fmt.Println("函数执行完成再调用")
  fmt.Println("func main()")
  test()
}
// defer栈
// 会先打印 9
func test()  {
  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }
}
