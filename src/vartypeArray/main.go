package main

import "fmt"

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
	// 数组的访问2，使用rank关键词
	for index, value := range arr1 {
		fmt.Println(index, value)
	}

}
