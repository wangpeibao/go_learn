package main

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

// 内置的strings库，各种操作

func main() {
	// EqualFold, 字符串是否相等，大小写不敏感
	fmt.Println(strings.EqualFold("abc", "Abc"))
	// HasPrefix, 是否以某个字符串为前缀
	fmt.Println(strings.HasPrefix("abc", "ab"))
	// HasSuffix, 是否以某个字符串未后缀
	fmt.Println(strings.HasSuffix("abc", "bc"))
	// Contains, 报否包含字符串, 其中空串和相等字符串也返回true
	fmt.Println(strings.Contains("abc", "abc"))
	// ContainsRune，字符串中是否包含utf8/unicode编码
	rune1 := []rune("我")
	fmt.Println(strings.ContainsRune("我爱中华", rune1[0]))
	// ContainsAny，两个字符串是否有重合的字符
	fmt.Println(strings.ContainsAny("abc", "cdef"))
	// Count返回包含子串的个数
	fmt.Println(strings.Count("abcdd", "d"))
	//　Index,返回子串第一次出现的位置，如果不存在则返回-1
	fmt.Println(strings.Index("abcd", "c"), strings.Index("abcd", "abd"))
	// IndexByte，返回byte第一次出现在字符串中的位置，没有则返回-1
	var byte1 byte = 97
	fmt.Println(strings.IndexByte("abcd", byte1))
	// IndexRune，返回rune第一次出现在字符串中的位置，没有则返回-1
	fmt.Println(strings.IndexRune("我爱中国", rune1[0]))
	// IndexAny,字符串chars中的任一utf-8码值在s中第一次出现的位置，如果不存在或者chars为空字符串则返回-1。
	fmt.Println(strings.IndexAny("abdc", "bdc"))
	// IndexFunc s中第一个满足函数f的位置i（该处的utf-8码值r满足f(r)==true），不存在则返回-1。
	f := func(r rune) bool {
		if r == rune1[0] {
			return true
		} else {
			return false
		}
	}
	fmt.Println(strings.IndexFunc("爱我中华", f))
	// LastIndex, LastIndexAny, LastIndexFunc, 参考上面的方法，只不过返回最后一次出现的位置
	// Title 返回字符串的首字母大写字符串
	fmt.Println(strings.Title("we have a new world"))
	// ToLower 将字符串中所有的都转为小写
	fmt.Println(strings.ToLower("AbCd"))
	// ToLowerSpecial 自定义大小到小写映射
	var delta = [3]rune{1, 2, 3} // 分别代表转大写事件，转小写事件，转标题事件对应的原值偏移量，如果希望转小写，原值＋delta[1]是对应的小写
	caseRange := unicode.CaseRange{
		Lo:    65,    // 范围的开始，包括这个值
		Hi:    65,    // 返回的结束, 包含这个值
		Delta: delta, // 三个时间对应的偏移
	}
	case1 := unicode.SpecialCase{caseRange}
	fmt.Println("AbCd", strings.ToLowerSpecial(case1, "AbCd"))
	// ToUpper将所有的转为大写
	fmt.Println(strings.ToUpper("aBcD abcd"))
	// ToTitle将所有字母都转为标题版本
	fmt.Println(strings.ToTitle("aBcD abcd"))
	// ToUpperSpecial 和 ToTitleSpecial 的功能实现和ToLowerSpecial相同
	// Repeat n个字符串的联结
	fmt.Println("3个abc", strings.Repeat("abc", 3))
	// Replace 替换旧的字符串，n表示替换前几个，如果为-1则表示全部替换
	fmt.Println(strings.Replace("aabbcc", "b", "B", 12))
	// Map 也是一个替换的操作，针对unicode编码，通过函数的形式进行替换
	f2 := func(r rune) rune {
		return r + 1
	}
	fmt.Println(strings.Map(f2, "abcd"))
	// Trim 相当于一个字符集删除操作，清楚cutset中所有在s一头一尾中的字符，中间的空格符不删除
	fmt.Println(strings.Trim("abbcd", "abef"))
	// TrimSpace 删除两端的空白字符，unicode.IsSpace包括'\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0
	fmt.Println(strings.TrimSpace("\r \nab cd\n \t"))
	// TrimFunc 删除两端符合方法返回true的字符
	f3 := func(r rune) bool {
		if r >= 'b' {
			return true
		} else {
			return false
		}
	}
	fmt.Println(strings.TrimFunc("abcd", f3))
	// 其中TrimLeft, TrimRight, TrimLeftFunc, TrimRightFunc都和Trim, TrimFunc功能相似，只不过时左端，右端
	// TrimPrefix, TrimSuffix，删除符合的前缀和后缀
	// Fields, 返回以一个多连续多个空白字符分割的字符串数组
	fmt.Println(strings.Fields("how old are you"))
	// FieldsFunc 使用函数来认定分割字符串
	f4 := func(r rune) bool {
		if r >= 'd' {
			return true
		} else {
			return false
		}
	}
	fmt.Println(strings.FieldsFunc("abcda aefbaw", f4))
	// Split 直接制定分隔符的分割,如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
	fmt.Println(strings.Split("a,b,,c", ","))
	// SplitN的功能和Split一样，只是多了一个返回切片的数量
	// SplitAfter 不会像Split把分隔符去掉而是在分隔符的后面分隔， SplitAfterN功能同上
	// Join 将字符串数组连接到一个字符串，有连接符
	a := []string{"123", "abc"}
	fmt.Println(a, strings.Join(a, ""))

	// 以下时Reader部分的
	// Reader时strings库里面定义的一个字符串读取数据类型，包括字符串，当前字符下标以及上一个读取字符的下标
	read := strings.NewReader("123456") // 创建一个Reader的指针数据类型
	fmt.Println("剩余字节数", read.Len())    // 返回Reader中未被读取的长度
	byte2, _ := read.ReadByte()         // 按字节读取字符串
	fmt.Println(byte2)
	fmt.Println("剩余字节数", read.Len())
	_ = read.UnreadByte() // 实际上时一个回退操作，回到上次读取的位置
	fmt.Println("剩余字节数", read.Len())
	// ReadRune和UnreadRune的功能实现同上，只是使用unicode读取
	// Seek, Reader查找位置的办法，偏移量和起始位置，并且将指针指到找到的位置
	index, err1 := read.Seek(2, io.SeekStart)
	fmt.Println(index, err1)
	index, err1 = read.Seek(2, io.SeekCurrent)
	fmt.Println(index, err1)
}
