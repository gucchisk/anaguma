package common

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

const versionFile = "cmd/version.go"

func SetVersion(version string) error {
	b, err := ioutil.ReadFile(versionFile)
	if err != nil {
		return err
	}
	lines := string(b)
	re := regexp.MustCompile(`version = .+`)
	lines = re.ReplaceAllString(lines, `version = "`+version+`"`)
	return ioutil.WriteFile(versionFile, []byte(lines), 0755)
}

func GetVersion() (string, error) {
	b, err := ioutil.ReadFile(versionFile)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`version = "(.*)"`)
	matches := re.FindSubmatch(b)
	if len(matches) < 2 {
		return "", fmt.Errorf("version not found")
	}
	return string(matches[1]), nil
}

func GetVersionNumber() (string, error) {
	v, err := GetVersion()
	if err != nil {
		return "", err
	}
	return v[1:], nil
}
