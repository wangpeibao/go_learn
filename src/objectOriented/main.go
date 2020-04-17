package main

import "fmt"

// golang的面向对象
type Account struct { // Account是一个类，其有属性phone和name
	phone string
	name  string
}

func (acc *Account) update(p, n string) (err error) { // 为Account类添加一个更新的方法,如果不需要修改就不适用*Account
	acc.phone = p
	acc.name = n
	return
}

func (acc Account) selfPrint() { // 顶一个一个Account的方法
	fmt.Println(acc)
}

// 定义一个继承Account的类
type Emp struct {
	acount Account
	icon   int
}

// 继承上述的方法
func (emp Emp) selfPrint() {
	emp.acount.selfPrint()
}

// 虚基类
type Test struct {
	s string
}

func (t Test) testprint() {
	fmt.Println(t.s)
}

type Test1 struct {
	*Test
	ss string
}

func main() {
	// go语言的面向对象与python的class(object)不太一样
	// go的实现方式时，用type + struct声明类与类的属性，用函数未类增加方法
	// 实例化一个类
	acc := new(Account) // 初始化一个空，struct中的值会有一个默认值,字符串为空字符串，数值是0, bool类型是false
	fmt.Println(acc)
	acc1 := &Account{} // 另一种初始化默认值的方法
	fmt.Println(acc1)
	acc2 := &Account{phone: "15620011759", name: "wang"} // 制定属性初始化
	fmt.Println(acc2)
	// 调用类的方法
	acc2.update("13512973681", "wangpeibao")
	fmt.Println(acc2)

	// 匿名组合，相当于继承
	emp := new(Emp)
	emp.selfPrint()
	// 还可以使用指针方式进行派生,
	t1 := &Test{"123"}
	t2 := &Test1{t1, "123"} // 这样就可以比较方便的使用Test的方法
	t2.testprint()
	// 当类型与其组合类都具有相同的属性名，哪个组合类中的属性会被隐藏
	// 可见性是通过属性的大小写和方法的大小写控制的

}
