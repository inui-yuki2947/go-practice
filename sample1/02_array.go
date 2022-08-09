package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----配列定義サンプル-----")
	var1 := [3]int{1, 2, 3}
	var2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(var1, var2)
	var1[0] = 0
	fmt.Println(var1, var2)

	fmt.Println("-----配列・スライスサンプル-----")
	var (
		var3 [3]int
		var4 []int
	)
	fmt.Println(var3, var4)
	var3[0] = 100
	var4 = append(var4, 100)
	fmt.Println(var3, var4)
}
