package cmd

import (
	"bytes"
	"testing"
)

func TestRootCmd(t *testing.T) {
	out := new(bytes.Buffer)
	rootCmd.SetOutput(out)
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("fail %v", err)
	}
	result := out.String()
	out.Reset()

	rootCmd.SetArgs([]string {"help"})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("fail %v", err)
	}
	result_help := out.String()
	out.Reset()

	rootCmd.SetArgs([]string {"--help"})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("fail %v", err)
	}
	result_help_flag := out.String()
	out.Reset()
	
	if result != result_help {
		t.Fatalf("not same result\ncmd:\n%s\nhelp:\n%s\n", result, result_help)
	}
	if result != result_help_flag {
		t.Fatalf("not same result\ncmd:\n%s\nhelp flag:\n%s\n", result, result_help_flag)
	}
}
