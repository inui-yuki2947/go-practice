package main

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func main() {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain("あああ", "いいい", false)
	fmt.Println(diffs)
}
