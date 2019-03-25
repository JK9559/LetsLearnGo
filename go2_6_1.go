package main

import (
	"fmt"
)

// 注意任意的类型都实现了空interface 也就是interface{}
type Human2_6 struct {
	name  string
	age   int
	phone string
}

type Student2_6 struct {
	Human2_6
	school string
	loan   float64
}

type Employee2_6 struct {
	Human2_6
	company string
	money   float64
}

// 定义interface
type Men2_6 interface {
	sayHi()
	sing(ly string)
	guzzle(beer string)
}

type YoungChap interface {
	sayHi()
	sing(ly string)
	borrowMoney(amount float64)
}

type ElderlyGent interface {
	sayHi()
	sing(ly string)
	spendMoney(amount float64)
}

// human实现sayHi方法
func (h *Human2_6) sayHi() {
	fmt.Printf("This is %s,u can call me on %s\n", h.name, h.phone)
}

// human实现sing方法
func (h *Human2_6) sing(ly string) {
	fmt.Printf("I can sing %s\n", ly)
}

// human实现guzzle方法
func (h *Human2_6) guzzle(br string) {
	fmt.Printf("Guzzle %s\n", br)
}

// emp重写human的sayhi
func (emp *Employee2_6) sayHi() {
	fmt.Printf("I word at %s\n", emp.company)
}

// student 实现 borrowmoney方法
func (stu *Student2_6) borrowMoney(amt float64) {
	stu.loan += amt
}

// employee 实现 spendmoney方法
func (emp *Employee2_6) spendMoney(amt float64) {
	emp.money -= amt
}
