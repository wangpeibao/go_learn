package main

// 接口功能
type Integer int // 定义一个Integer的类
func (a *Integer) AddOne() { // 实现一个+1的操作
	*a += 1
}
func (a Integer) AddB(b Integer) Integer {
	return a + b
}

type Adder interface { // 顶一个累加1接口
	AddOne()
	AddB(b Integer)
}

func main() {
	// 一个类只要实现了接口所要求的所有函数，就说这个类实现了该接口。也就可以将这个类的实例赋值给该接口

}
