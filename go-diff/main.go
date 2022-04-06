package main

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
)

func main() {
	// 単純な比較
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain("あああ", "いいい", false)
	fmt.Println(diffs)

	// ファイルを読み込んで比較
	text1, _ := ioutil.ReadFile("text/text1.txt")
	text2, _ := ioutil.ReadFile("text/text2.txt")
	diffs = dmp.DiffMain(string(text1), string(text2), false)
	fmt.Println("--------------------------------------------------")
	fmt.Println(diffs)
}
