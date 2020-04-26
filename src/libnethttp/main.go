package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// golang提供了http相关的标准库

func main() {
	// func Get http的get请求, 返回http.Response定义的结构体
	res, err := http.Get("http://www.baidu.com")
	if err == nil {
		fmt.Println(res.StatusCode) // 状态码，例如200,404等
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body) // 返回结构体的内容
		if err == nil {
			fmt.Println(reflect.TypeOf(body).String())
		}
	} else {
		fmt.Println(err) // 如果访问不成功会返回错误信息
	}

}
