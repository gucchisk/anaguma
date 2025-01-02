package main

import (
	"fmt"

	"github.com/gucchisk/anaguma/common"
)

func main() {
	v, err := common.GetVersion()
	if err != nil {
		return
	}
	fmt.Print(v)
}
