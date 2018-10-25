package main

//import中并没有逗号分隔
//可以编写多个导入语句，但推荐使用打包的导入语句。
import (
  "fmt"
  "math/rand"
)

// 字符串连接不需要连接符
func main()  {
  fmt.Println("My favorite number is ",rand.Intn(1000))
}

// 在go中，首字母大写的名称是被导出的，如下出现一下错误，要检查首字母大小写。
// annot refer to unexported name math.pi
