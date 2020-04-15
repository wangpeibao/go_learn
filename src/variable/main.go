package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 标准的声明变量, var variable_name variable_type = value;其中value的数据类型必须和variable_type的类型一致
	var x int = 10
	fmt.Println(x)
	// go语言支持类型推导，有点像python，但是它仍然是一个静态数据类型语言
	var y = 10
	fmt.Println(reflect.TypeOf(y).String())
	// 这个时候y的数据类型已经是int数据类型了，当我把y赋值成一个字符串时会报错
	// go的另一种简易的数据声明初始化, 和上面的生命效果是一样的
	z := 10
	fmt.Println(reflect.TypeOf(z).String())
	// 对于同一个变量不能重复声明
	// go语言引入了多重赋值，操作和python一样
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a, b)
	// 输出的结果是a,b的值进行了交换

	// 匿名变量，当我们调用函数时，有些返回的值时不需要的，就可以使用匿名函数进行占位形如: _, c := func1(),其中 _ 就是匿名变量
}
