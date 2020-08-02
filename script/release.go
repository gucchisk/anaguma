package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
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

	versionFile := "cmd/version.go"

	b, err := ioutil.ReadFile(versionFile)
	if err != nil {
		panic(err)
	}
	lines := string(b)
	re := regexp.MustCompile(`version = .+`)
	lines = re.ReplaceAllString(lines, `version = "` + versionStr + `"`)
	err = ioutil.WriteFile(versionFile, []byte(lines), 0755)
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
