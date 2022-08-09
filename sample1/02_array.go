package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----配列定義サンプル-----")
	b1 := [3]int{1, 2, 3}
	b2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b1, b2)
	b1[0] = 0
	fmt.Println(b1, b2)

	fmt.Println("-----配列・スライスサンプル-----")
	var (
		d1 [3]int
		d2 []int
	)
	fmt.Println(d1, d2)
	d1[0] = 100
	d2 = append(d2, 100)
	fmt.Println(d1, d2)
}
