package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("-----変数定義サンプル-----")
	var a1 string
	a1 = "A1"
	a2 := "A2"
	var (
		a3 = 1
		a4 = 2
	)
	fmt.Println(a1, a2, a3, a4)

	fmt.Println("-----配列定義サンプル-----")
	b1 := [3]int{1, 2, 3}
	b2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b1, b2)
	b1[0] = 0
	fmt.Println(b1, b2)

	fmt.Println("-----関数定義サンプル-----")
	var c3 func(int, string) int
	fmt.Println(c1(2, 3))
	fmt.Println(c2)
	fmt.Println(c2())
	fmt.Println(c3)

	fmt.Println("-----配列・スライスサンプル-----")
	var (
		d1 [3]int
		d2 []int
	)
	fmt.Println(d1, d2)
	d1[0] = 100
	d2 = append(d2, 100)
	fmt.Println(d1, d2)

	fmt.Println("-----外部ライブラリサンプル-----")
	app := cli.NewApp()
	fmt.Println(app)
}

func c1(x, y int) int {
	return x + y
}

func c2() (x int) {
	x = 1
	return
}
