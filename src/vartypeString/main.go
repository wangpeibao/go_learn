package main

import (
	"fmt"
)

func main() {
	// 字符串类型 声明关键字 string
	var str1 string = "abcd"
	str2 := "1234"
	fmt.Println(str1, str2)
	// 可以用下标的方式获取字符串中的内容，但是初始化后不可修改
	fmt.Println("取字符串第一第二位", str1[0], str1[1])
	// 使用内置的len方法获取字符串的长度,这里的长度是字节长度，出现中文时，一个中文占三个字节
	str3 := "123你好"
	fmt.Println("字符串长度字节", str3, len(str3))
	// 用+进行字符串的连接
	fmt.Println("str1 + str2 =", str1+str2)
	// 字符串遍历1,使用下标的方式遍历获取,这是获取的字节
	for i := 0; i < len(str3); i++ {
		fmt.Println("以字节依次输出字符串中的字符", i, str3[i])
	}
	// 字符串遍历2,使用range方法，以unicode遍历,ch的数据类型是rune
	for i, ch := range str3 {
		fmt.Println("以unicode方式的遍历", i, ch)
	}
	// 字符串切片，和python的切片基本相同, 是一个下标的左闭右开区间，左端点如果是0可以省略，右端点如果是末尾，也可以省略,不支持倒序负数,不允许越界
	fmt.Println(str3)
	fmt.Println("下标的左闭右开区间str3[1:3] =", str3[1:])
	fmt.Println("省略左端点，从头开始str3[:3] =", str3[:3])
	fmt.Println("省略右端点，直到末尾str3[1:] =", str3[1:])
	// 字符串初始化不支持修改，可转化成byte数组或者rune数组
	byte1 := []byte(str3)
	fmt.Println(byte1)

	// 字符类型 byte单字节, rune单个unicode字符两种类型
}
