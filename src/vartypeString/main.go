package main

import "fmt"

func main() {
	// 字符串类型 声明关键字 string
	var str1 string = "abcd"
	str2 := "1234"
	fmt.Println(str1, str2)
	// 可以用下标的方式获取字符串中的内容，但是初始化后不可修改
	fmt.Println("取字符串第一第二位", str1[0], str1[1])
}
