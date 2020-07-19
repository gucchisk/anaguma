/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"github.com/gucchisk/bytestring"
	"github.com/spf13/cobra"
)

var out string
var outputFormat Format

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "anaguma",
	Short: "badger DB clinet utilities",
	Long: "anaguma is a CLI tool to access badger DB",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initialize)
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&out, "out", "o", "ascii", "output format [ascii hex base64]")
}

func initialize() {
}

type Format int
const (
	Unknown Format = iota
	Ascii
	Hex
	Base64
)

func NewFormat(format string) (Format, error) {
	switch format {
	case "ascii":
		return Ascii, nil
	case "hex":
		return Hex, nil
	case "base64":
		return Base64, nil
	default:
		return Unknown, fmt.Errorf("unknown format: %s", format)
	}
}

func byteToStr(b []byte, out Format) string {
	bytes := bytestring.NewBytes(b)
	switch out {
	case Ascii:
		return bytes.String()
	case Hex:
		return bytes.HexString()
	case Base64:
		return bytes.Base64()
	default:
		return ""
	}
}

func strToByte(s string, in Format) ([]byte, error) {
	var b bytestring.Bytes
	var err error
	switch in {
	case Ascii:
		b, err = bytestring.NewBytesFrom(s, bytestring.Normal)
	case Hex:
		b, err = bytestring.NewBytesFrom(s, bytestring.Hex)
	case Base64:
		b, err = bytestring.NewBytesFrom(s, bytestring.Base64)
	default:
		return []byte{}, fmt.Errorf("unknown format")
	}
	return b.ByteArray(), err
}
