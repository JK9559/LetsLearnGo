package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type human struct {
	name   string
	age    int
	weight int
}

type student struct {
	human
	speciality string
}

func older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func structTest() {
	var tom person

	tom.name, tom.age = "Tom", 8

	bob := person{name: "Bob", age: 18}

	paul := person{"Paul", 28}

	tb_Older, tb_diff := older(tom, bob)
	tp_Older, tp_diff := older(tom, paul)
	bp_Older, bp_diff := older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bp_Older.name, bp_diff)

	mark := student{human{"Mark", 24, 48}, "ahahaha"}
	fmt.Println("His name is: ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

	mark.speciality = "AI"
	mark.age = 46
	mark.weight += 60
	fmt.Println("His name is: ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
}
