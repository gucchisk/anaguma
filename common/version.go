package common

import (
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
	lines = re.ReplaceAllString(lines, `version = "` + version + `"`)
	return ioutil.WriteFile(versionFile, []byte(lines), 0755)
}
