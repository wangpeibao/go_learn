package main

import "fmt"

// 函数功能
func AddOne(x int) int { // 当有一个返回值时
	return x + 1
}

func Add(x, y int) (int, bool) { // 当有个多个返回值时
	return x + y, true
}

func AddAll(nums ...int) int {
	count := 0
	for _, num := range nums {
		count += num
	}
	return count
}

func main() {
	// 函数的定义，使用func关键字,形如上AddOne和Add方法
	a := AddOne(1)
	fmt.Println(a)
	b, status := Add(1, 2)
	fmt.Println(b, status)
	// 函数调用，大写才在可以在本package外使用，小写只可以package内使用
	// 不定参数，首先应该让函数可以接受不定参数，形如AddAll方法，但是不定参数只能放在所有参数的最后一个
	// 不定参数的传递如下，不定参数也可以设置成接受各种数据类型，使用interface{}接口
	c := AddAll(1, 2, 3, 4, 5) // 正常传递
	fmt.Println(c)
	arr1 := []int{1, 2, 3, 4, 5}
	c = AddAll(arr1...) // 切片解析方式传递
	fmt.Println(c)
	// 多返回值，go语言中可以为返回值命名，会为命名的参数初始化值，在return不注明返回的数据，也会把参数返回
	// 匿名函数，没有函数名的函数，可用于给参数赋值和直接使用
	fmt.Println(func(x int) int { return x + 1 }(1)) // 适用与较为简单的

}
