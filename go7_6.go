package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串操作
func strOp() {
	// 字符串 是否包含子串 返回bool值
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println("=========")

	// 字符串连接 把[]string 用 另一个string连接起来
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, "^-^"))
	fmt.Println("=========")

	// 字符串中查找子串的位置 查找不到返回-1
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	fmt.Println("=========")

	// 字符串重复多少次 返回重复的字符串
	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println("=========")

	// 在字符串中，把old字符串替换为new，n为替换次数，小于0为全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	fmt.Println("=========")

	// 把字符串按分隔符分隔 返回string slice
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	fmt.Println("=========")

	// 在字符串头尾去掉指定的字符串
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", " !A"))
	fmt.Println("=========")

	// 去掉字符串的空格符 并按照空格分隔slice
	fmt.Printf("|%q|\n", strings.Fields("    foo bar     baz     "))
}

// Append系列函数 将整数等 转化为字符串之后，添加到现有的字符数组中
func strAppend() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '啥')
	fmt.Println(string(str))
	// 4567false"abcdefg"'啥'
}

// Format系列函数 将不同的类型装换为字符串
func strFmt() {
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

// Parse系列函数把字符串转化为其他类型
func strParse() {
	a, err := strconv.ParseBool("false")
	checkError(err)

	b, err := strconv.ParseFloat("123.234", 64)
	checkError(err)

	c, err := strconv.ParseInt("12334", 10, 64)
	checkError(err)

	d, err := strconv.ParseUint("1234", 10, 64)
	checkError(err)

	e, err := strconv.Atoi("1023")
	checkError(err)

	fmt.Println(a, b, c, d, e)
}

func strTrans() {
	strAppend()
	strFmt()
	strParse()
}

func go7_6() {
	strTrans()
}
