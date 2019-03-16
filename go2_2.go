package main

import (
	"fmt"
)

const (
	PI   = 3.14
	PIPI = 3.1415926
)

func valTest() {
	var a int

	var b, c, d int

	var e int = 0

	var f, g, h float32 = 1.0, 2.0, 3.0

	var i, j, k = 1, 2.0, "varK"

	l, m, n := "cc", 1, 3.0

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, PI, PIPI)

}
