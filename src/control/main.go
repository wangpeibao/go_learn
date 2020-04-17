package main

import "fmt"

// 控制语句，包括循环，条件，跳转
func testLoop(x int) int {
	if x == 1 {
		return 1
	} else if x == 2 {
		return 2
	} else {
		return x
	}
}

func testSwitch(x int) {
	switch x {
	case 1:
		{
			fmt.Println("就是１")
		}
	case 2:
		{
			fmt.Println("就是2")
		}
	default:
		{
			fmt.Println("其他")
		}
	}
}

func main() {
	// 条件语句 支持if ... else if ... else ...
	a := 1
	fmt.Println(testLoop(a))

	// 选择语句 switch case, 如果不指定switch的条件，实现和case相同
	b := 2
	testSwitch(b)

	// 循环语句, 只支持for关键字，当没有三段条件时，认为是死循环
	for i := 0; i < 5; i++ {
		fmt.Println("1")
	}
	i := 1
	for {
		i += 1
		fmt.Println(i)
		if i >= 10 {
			break
		}
	}

	// 跳转语句　goto关键字
}
