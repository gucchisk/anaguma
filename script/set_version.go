package main

import (
	"os"
	"regexp"
	"github.com/gucchisk/anaguma/common"
)

func main() {
	versionStr := os.Args[1]
	ok, err := regexp.MatchString(`v\d+\.\d+\.\d+`, versionStr)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("invalid version (value: " + versionStr + ")")
	}
	err = common.SetVersion(versionStr)
	if err != nil {
		panic(err)
	}
}
