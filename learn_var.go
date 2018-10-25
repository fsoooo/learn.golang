package main

import (
  "fmt"
)
//变量定义 关键字/变量名/变量类型/初始值
// var num int
// var num int = 0 等价
var number_value int = 20
var string_value string = "hello world!!!"
var bool_vulue bool = true
var num1,num2 int = 1,2
func main()  {
  fmt.Println(number_value,string_value,bool_vulue)
  fmt.Println(num1,num2)
  test()
}

// 短声明变量 :=  修饰明确类型的变量，来替代var，只能在函数内使用
func test()  {
  a := 5
  fmt.Println(a)
}

// basic types
// bool string int系 uint系 byte rune float系 complex系

// 常量 const
