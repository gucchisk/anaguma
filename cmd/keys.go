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
	"errors"
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/gucchisk/anaguma/common"
	"github.com/gucchisk/anaguma/db"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys <dir>",
	Short: "Show keys",
	Long: "Show keys",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("require badger DB dir")
		}
		var err error
		outputFormat, err = common.NewFormat(out)
		if err != nil {
			return err
		}
		return nil;
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := db.Keys(args[0], func(keys [][]byte) {
			for _, v := range keys {
				value := common.ByteToStr(v, outputFormat)
				fmt.Printf("%v\n", value)
			}
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keysCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keysCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
