package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----構造体サンプル-----")
	fmt.Println(Hoge{1, "aaa"})
	// fmt.Println(Hoge{1}) コンパイルエラー
	fmt.Println(Hoge{Hoge1: 1, Hoge2: "aaa"})
	fmt.Println(Hoge{Hoge1: 1})
	fmt.Println(Hoge{})
}

type Hoge struct {
	Hoge1 int
	Hoge2 string
}
