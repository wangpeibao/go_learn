package main

import (
	"fmt"
	"reflect"
)

// 数据类型

func main() {
	// 布尔类型，用bool声明，对应着true, false
	var bool_1 bool = true
	bool_2 := false
	fmt.Println("布尔类型的声明初始化", bool_1, bool_2)
	// 布尔类型不支持强转和其他类型的赋值bool3 := 1 和 bool4 = bool(1)都不对，支持条件判断推出
	bool_3 := 1 == 1
	fmt.Println("条件推到出", bool_3)

	// 整型类型，包括int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintpr
	// int*表示有符号数，uint*表示无符号数
	// int, uint, uintpr都和平台相关，但是即使在32位操作系统中，int和int32即使表示的数据的范围时一样的，也时不同的数据类型，进行赋值时会报错
	// 可使用强制转换，注意由大到小和由浮点到整数的精度丢失
	var a int32 = 1
	fmt.Println("定义为int32的数值a", a, reflect.TypeOf(a).String())
	var b int = int(a)
	fmt.Println("用强转之后的a来赋值int类型的b", b, reflect.TypeOf(b).String())
	// 整型的数值运算
	// 运算操作 基本的+-*/%运算
	// 比较运算 == != > < >= <=
	// 位运算 左移:<<, 右移:>>, 与:&, 或:|, 异或:^, 取反:^

	// 浮点数, 包括float32和float64,时单精度和双精度浮点数，不声明变量类型时，默认float64
	float_1 := 1.0  // 如果没有小数点会认为时int型
	fmt.Println("默认float类型", reflect.TypeOf(float_1).String())
	// 可以使用强制转化， float32->float64, float->int
	float_2 := 2.5
	floatInt := int(float_2)
	fmt.Println("强制float->int", floatInt, reflect.TypeOf(floatInt))
	// 因为时静态数据类型，也不支持float与int之间的+-
	// float类型的比较，可以是math.Pdim(f1, f2) < p来处理， 其中p时精度，形如0.00001

	// 复数类型, complex64(两个float32表示实部和虚部)和complex128(两个float64表示实部和虚部)
	var com complex64 = 1 + 2i
	com1 := 1 + 3i
	com2 := complex(1, 4)
	fmt.Println("复数", com, com1, com2)

}
