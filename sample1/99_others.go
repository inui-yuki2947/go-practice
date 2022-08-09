package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func main() {
	fmt.Println("-----外部ライブラリサンプル-----")
	app := cli.NewApp()
	fmt.Println(app)
}
