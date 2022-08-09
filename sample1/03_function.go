package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----関数定義サンプル-----")
	var c3 func(int, string) int
	fmt.Println(c1(2, 3))
	fmt.Println(c2)
	fmt.Println(c2())
	fmt.Println(c3)
}

func c1(x, y int) int {
	return x + y
}

func c2() (x int) {
	x = 1
	return
}
