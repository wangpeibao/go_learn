package main

import (
	"fmt"
	"reflect"
)

// 数组和数组切片

func main() {
	// 数组的声明, 形如[length]type length为常量或常量表达式，一旦声明，数组的长度不可更爱，type时数组内元素的数据类型
	arr1 := []int{1, 2}
	fmt.Println(arr1)
	// 多维数据的声明和赋值
	arr2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr2)
	fmt.Println("使用len获取元素长度", len(arr2))
	// 数组的访问1,使用下标
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	// 数组的访问2，使用range关键词
	for index, value := range arr1 {
		fmt.Println(index, value)
	}

	// 数组是值类型，每次传递都产生副本
	// 数组切片就像一个指向数组的指针。
	// 数组切片的创建1,基于数组的实现,切片规则下标左闭右开，　左0和右length可以忽略
	arr3 := arr1[:1]
	fmt.Println(reflect.TypeOf(arr3).String())
	// 数组切片的穿件2,基于内置的make方法
	arr4 := make([]int, 1, 2) // 初始元素１个，初始元素值是0,预留2个存储控件
	arr5 := []int{1, 2, 3}    // 初始化长度为3的切片
	fmt.Println(arr4, arr5)
	// 元素遍历 和数组一样，可以适用下标，也可以适用range方法
	// 动态增加元素, 适用内置的append方法
	arr5 = append(arr5, 4, 5) // 添加元素
	fmt.Println(arr5)
	arr5 = append(arr5, arr4...) // 添加另一个切片，注意那三个点
	fmt.Println(arr5)
	// 在切片的基础上可以继续切片，与在数组上切片的不同是，可以超出长度的范围,但是小于分配的部分，超出的部分是默认值
	arr6 := []int{1, 2, 3, 4}
	fmt.Println(cap(arr6))
	// 切片的内容复制，适用copy函数, 如果长复制给短，那么只会复制短的那部分给短的，　如果短的复制给长的，只会复制长的前n个
	arr7 := []int{1, 2}
	arr8 := []int{3, 4, 5}
	copy(arr7, arr8)
	fmt.Println(arr7)
}
