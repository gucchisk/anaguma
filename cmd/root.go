/*
Copyright © 2020 gucchisk <gucchi_sk@yahoo.co.jp>

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
	"github.com/spf13/cobra"
	"github.com/gucchisk/anaguma/common"
	db1 "github.com/gucchisk/anaguma/db/v1"
	db2 "github.com/gucchisk/anaguma/db/v2"
)

var in string
var inputFormat common.Format
var out string
var outputFormat common.Format
var badgerVersion uint8
var db common.DB
var verbose bool
var showVersion bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "anaguma",
	Short: "badger DB clinet utilities",
	Long: "anaguma is a CLI tool to access badger DB",
	// Version: "v0.0.2",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// },
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		outputFormat, err = common.NewFormat(out)
		if err != nil {
			return err
		}
		switch badgerVersion {
		case 1:
			db = &db1.DB{}
		case 2:
			db = &db2.DB{}
		}
		return nil
	},
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
	rootCmd.PersistentFlags().StringVarP(&out, "out", "o", "ascii", "output format [ascii hex base64]")
	rootCmd.PersistentFlags().Uint8VarP(&badgerVersion, "bversion", "b", 1, "badger version [1, 2]")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "log", "l", false, "log output")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show anaguma version")
	rootCmd.Version = version
}

func initialize() {
}
