package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 常量时指编译不可改变的值，包括数值型，布尔型，字符型

	// 字面常量，硬编码常量，具体的常量类型，如true, 1, "123"等

	// 常量的声明和赋值 const constant_name constant_type = value
	// constant_type可选
	const a int = 10
	fmt.Println(a, reflect.TypeOf(a).String())
	const b = 10
	fmt.Println(b, reflect.TypeOf(b).String())
	// value可以时字面常量，也可以是在编译期运算的常量表达式，例如
	const c int = 10 + 10
	fmt.Println(c)
	// value不可以在运行时才可以得到的值

	// 预定义常量 iota, 在const关键字出现后被重置为0,之后每出现一次，自增1，知道下次const出现再重置为0
	const e = iota
	const f = iota
	fmt.Println(e, f)
	const (
		h = iota
		i
		j
		k
	)
	fmt.Println("当定义一组参数时依次递增", h, i, j, k)
	// go 没有enum枚举关键字，可以使用iota参数实现，形如上面hijk的定义

}
