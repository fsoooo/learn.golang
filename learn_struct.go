package main

import "fmt"

/**
结构struct
Go中的struct与C中的struct非常相似，并且Go没有class，代替了class的位置，但并没有代替class的功能
使用type
支持指向自身的指针类型成员
支持匿名结构，可用作成员或定义成员变量
匿名结构也可以用于map的值
可以使用字面值对结构进行初始化
允许直接通过指针来读写结构成员
相同类型的成员可进行直接拷贝赋值
支持==与!=比较运算符，但不支持>或<
支持匿名字段，本质上是定义了以某个类型名为名称的字段
嵌入结构作为匿名字段看起来像继承，但不是继承
可以使用匿名字段指针
 */
type Student struct {
  name string
  sex int
}

func main()  {
  student := Student{"xiaoming",1}
  fmt.Println(student.name)
  fmt.Println(student.sex)
}
