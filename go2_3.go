package main

import (
	"fmt"
)

func funTest() {
	// if语句允许在条件判断语句里声明一个变量，这个变量的作用域只能在该条件逻辑块内
	// 其他地方不起作用
	x := 2
	if x > 3 {
		fmt.Printf("valis: %d\n", x)
	} else if y := returnVal(); y >= 5 {
		fmt.Printf("valis: %d\n", y)
	} else {
		fmt.Println("aah")
	}

	// goto语句
	// goto语句大小写敏感
	// i := 0
	// Here:
	// fmt.Println(i)
	// i++
	// goto Here

	// defer 是延迟语句，可以在函数中添加多个defer语句，当函数执行到最后，
	// defer语句会按照逆序执行，最后该函数返回
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	// Panic & Recover

}

func returnVal() int {
	return 5
}

// 函数作为值、类型传递
// 声明函数类型 定义类型即可 不需要在入参和返回值添加参数名
type funcType func(int, float32) int

func yesICan(a int, b float32) int {
	if a < 10 {
		return 0
	}
	if b > 2.0 {
		return 1
	}
	return 1
}

func noUCant(a int, b float32) int {
	if a < 20 {
		return 2
	}
	if b > 3.0 {
		return 3
	}
	return 1
}

func testFunc(a int, f funcType) int {
	var b float32 = 1.0
	return f(a, b)
}
