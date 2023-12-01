package main

import (
	"os"
	"os/exec"
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

	cmd := exec.Command("git", "commit", "-am", "release " + versionStr)
	cmd.Start()
	cmd.Wait()
	cmd = exec.Command("git", "tag", versionStr)
	cmd.Start()
	cmd.Wait()
	cmd = exec.Command("git", "push", "origin", "master", "--tags")
	cmd.Start()
	cmd.Wait()
}
