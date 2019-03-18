package main

import (
	"fmt"
)

// 通过const定义常量
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

// Color作为byte的别名
type Color byte

// 定义Box长宽高 颜色属性
type Box struct {
	width, height, depth float64
	color                Color
}

// 定义slice，内含Box
type BoxList []Box

// 定义了接受者为Box，计算体积
func (b Box) volume() float64 {
	return b.width * b.depth * b.height
}

// 更改box的颜色 如果为传值不加指针，无法修改成功
// 原因为 如果不传Box的指针，那么函数接收的是Box的一个copy
// 也就是对于颜色值的修改 实际上只是作用于了Box的copy，并不是真正的Box，所以需要传入指针
func (b *Box) setColor(c Color) {
	b.color = c
}

// 返回体积最大的box的颜色
func (bList BoxList) biggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, box := range bList {
		if bVol := box.volume(); bVol > v {
			v = bVol
			k = box.color
		}
	}
	return k
}

func (bList BoxList) paintItBlack() {
	for i, _ := range bList {
		bList[i].setColor(BLACK)
	}
}

func (color Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[color]
}

// method 继承
type People struct {
	name  string
	age   int
	phone string
}

type stuPeople struct {
	People //匿名字段
	school string
}

type empPeople struct {
	People
	company string
}

// People 定义了自己的method
func (p *People) sayHi() {
	fmt.Printf("Hi ,I am %s\n", p.name)
}

// stuPeople的method重写了People的
func (stu *stuPeople) sayHi() {
	fmt.Printf("Hello, I come from %s\n", stu.school)
}

// empPeople的method重写了People的
func (emp *empPeople) sayHi() {
	fmt.Printf("Hello, I am an emp from %s\n", emp.company)
}

// method继承
// 如果匿名字段实现了一个method 那么包含这个匿名字段的struct也可以调用这个method
func testMethodSucc() {
	mark := stuPeople{People{"Mark", 25, "010-0099"}, "ABC Sch"}
	kitty := empPeople{People{"Kitty", 33, "861123"}, "Google"}

	mark.sayHi()
	kitty.sayHi()
}

// method 重写
func testMethodOver() {
	mark := stuPeople{People{"Mark", 25, "010-0099"}, "ABC Sch"}
	kitty := empPeople{People{"Kitty", 33, "861123"}, "Google"}

	mark.sayHi()
	kitty.sayHi()
}

func testMethod() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		// 注意这里最后一行也是有逗号的
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.biggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.paintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.biggestColor().String())

}
func test2_5() {
	testMethod()
	testMethodSucc()
	testMethodOver()
}
