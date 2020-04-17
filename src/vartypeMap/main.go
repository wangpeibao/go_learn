package main

import "fmt"

// map映射，字典类型

func main() {
	// map的声明, 如果只声明未初始化又没有创建的，无法赋值
	map1 := make(map[string]int) // 使用make创建
	map1["1"] = 1
	fmt.Println(map1["1"])
	map2 := map[string]int{
		"2": 2,
	}
	fmt.Println(map2["2"])
	// map的删除，适用内建delete方法
	map1["2"] = 2
	fmt.Println(map1)
	delete(map1, "2")
	fmt.Println(map1)
	// map的取值, 除了要查询的结果，还返回了一个是否查到的bool状态
	value, ok := map1["1"]
	fmt.Println(value, ok)
	value, ok = map1["2"]
	fmt.Println(value, ok)
}
