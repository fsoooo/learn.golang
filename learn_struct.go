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
	name  string
	sex   int
	age   int
	phone int
}
type T struct {
	Name string
	Age  int
}
/**
在go语言中，定义struct的时候，
除了类型和字段名一一对应，也可以支持不写字段名，
而只写类型，也就是匿名字段。
*/
//匿名字段
type Cat struct {
	int
	string
}


func main() {
	student := Student{"xiaoming", 1, 18, 15701681524}
	fmt.Println(student.name)
	fmt.Println(student.sex)
	fmt.Println(student.age)
	fmt.Println(student.phone)

	t := T{}
	fmt.Println(t)
	t.Name = "astar"
	t.Age = 10
	fmt.Println(t)
	t.Add()

	cat := Cat{12,"小花"}
	fmt.Println(cat.string,cat.int);
}
/**
新增了一个变种函数(其实是方法)，(t *T) 这就是给这个结构体绑定函数，
然后在结构体中就可以直接调用Add这个方法
GO就是以这种形式来实现面像对象的思想
如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
类型别名不会拥有底层类型所附带的方法
方法可以调用结构中的非公开字段
 */
func (t *T) Add() {
	fmt.Println(t.Age, t.Name)
}
