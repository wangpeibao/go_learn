package main

import "fmt"

// 接口功能
type Car interface {
	GetName() string
	Run()
}

type BMW struct {
	name string
}

func (c BMW) GetName() string {
	return c.name
}

func (c BMW) Run() {
	fmt.Printf("%s is running\r\n", c.name)
}

type BYD struct {
	name string
}

func (c BYD) GetName() string {
	return c.name
}

func (c BYD) Run() {
	fmt.Printf("%s is running\n\r", c.name)
}

func main() {
	// 一个类只要实现了接口所要求的所有函数，就说这个类实现了该接口。也就可以将这个类的实例赋值给该接口
	byd := BYD{"BYD"}
	bmw := BMW{"BMW"}
	var c Car
	c = byd
	c.Run()
	c = bmw
	c.GetName()
	c.Run()
	// 接口允许嵌套实现

	// 对与从接口获取实际的类型可以有多种实现方式
	// 直接对于接口断言
	var b interface{}
	d := 64
	b = d
	value, ok := b.(int) // 对接口b进行int的断言，如果b是int, value是断言之后的数据类型，ok成功与否
	fmt.Println(value, ok)
	// 断言接口的类型，通过switch,case获取
	switch b.(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("other")
	}
}
