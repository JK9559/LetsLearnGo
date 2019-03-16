package main

import (
	"fmt"
)

const (
	PI   = 3.14
	PIPI = 3.1415926
)

func valTest() {

	// 变量的声明
	var a int

	var b, c, d int

	var e int = 0

	var f, g, h float32 = 1.0, 2.0, 3.0

	var i, j, k = 1, 2.0, "varK"

	l, m, n := "cc", 1, 3.0

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, PI, PIPI)

	// 数值类型
	// 不同类型的变量不能相互赋值
	// 复数
	var compl complex64 = 5 + 5i
	fmt.Printf("val is : %v\n", compl)

	// 字符串时不可变的，如果需要可变可按照下面方法
	str := "hello"
	cstr := []byte(str)
	cstr[0] = 'c'
	str2 := string(cstr)
	fmt.Printf("ori :%s, now :%s\n", str, str2)
	// 也可以以下方法 切片
	str3 := "c" + str[1:]
	fmt.Printf("now1 :%s\n", str3)

	// map 无序 长度不定 可改 不是线程安全的
	rating := map[string]float32{"C": 5, "Go": 4.5, "C++": 5.0}
	// map 有两个返回值，第二个返回值，如果key不存在为false，存在未true

	_, ok := rating["C#"]
	if ok {
		fmt.Println("Yes We have")
	} else {
		fmt.Println("No We haven`t")
	}

	// map 也是引用类型 如果两个map指向同一个底层，一个改变，另一个也会改变
	mMap := make(map[string]string)
	mMap["Hello"] = "Hahaa"
	mMap1 := mMap
	mMap1["Hello"] = "abc"
	fmt.Println(mMap, mMap1)

	// 函数new和函数make
	// 函数new分配了零值填充的类型T的内存空间，并且指向了其地址，*T
	// 函数make只能创建 slice,map,channel 返回其T类型 不是*T
	// make初始化时 初始化了其数据结构 cap len之类的
	// new 返回指针
	// make 返回初始化后的非零值
}
