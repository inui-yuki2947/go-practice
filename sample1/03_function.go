package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----関数定義サンプル-----")
	fmt.Println(function1(2, 3))
	fmt.Println(function2)
	fmt.Println(function2())

	var function3 func(int, string) int
	fmt.Println(function3)
}

func function1(x, y int) int {
	return x + y
}

func function2() (x int) {
	x = 1
	return
}
