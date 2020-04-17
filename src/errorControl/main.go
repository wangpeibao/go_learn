package main

import (
	"fmt"
	"reflect"
)

// 错误处理
type MyError struct {
	err string
}

func (e *MyError) Error() string {
	return e.err
}

func getError() (err error) {
	return &MyError{"这是一条错误的日志"}
}

func main() {
	// error接口, 适用error接口，自定义错误日志
	a := getError()
	fmt.Println(reflect.TypeOf(a).String())

	// defer关键字，当函数退出时，会执行defer后面的命令,用于释放资源，如果比较复杂，也可以适用匿名函数

	// panic,recover关键字
	// panic,当调用panic时，正常的函数流程会被终止，但是不影响defer的执行
	// recover用于终止错误处理，流程接待panic的参数
}
