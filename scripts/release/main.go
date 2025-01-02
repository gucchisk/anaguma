package main

import (
	"fmt"
	"os"

	"github.com/gucchisk/anaguma/common"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: release v<version>")
		return
	}
	v := os.Args[1]
	common.SetVersion(v)
	fmt.Printf("set version %s\n", v)
}
