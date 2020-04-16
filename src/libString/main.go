package main

import (
	"fmt"
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

}
