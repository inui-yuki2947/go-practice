package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("-----変数定義サンプル-----")
	var var1 string
	var1 = "A1"
	var2 := "A2"
	var (
		var3 = 1
		var4 = 2
	)
	fmt.Println(var1, var2, var3, var4)
}
