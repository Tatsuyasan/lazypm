package main

import (
	"fmt"
	"os"

	"github.com/Tatsuyasan/lazyPm/packages/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
