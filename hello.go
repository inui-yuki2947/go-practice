package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// 変数定義サンプル
	var a1 string
	a1 = "A1"
	a2 := "A2"
	var (
		a3 = 1
		a4 = 2
	)
	fmt.Println(a1, a2, a3, a4)

	// 配列定義サンプル
	b1 := [3]int{1, 2, 3}
	b2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b1, b2)
	b1[0] = 0
	fmt.Println(b1, b2)
}
