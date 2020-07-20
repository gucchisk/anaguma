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
)

// valuesCmd represents the values command
var valuesCmd = &cobra.Command{
	Use:   "values <dir>",
	Short: "Show keys & values",
	Long: "Show keys & values",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("require badger DB dir")
		}
		return nil;
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := db.Open(args[0], verbose)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.Values(func(keys, values [][]byte) {
			for i, k := range keys {
				key := common.ByteToStr(k, outputFormat)
				value := common.ByteToStr(values[i], outputFormat)
				fmt.Printf("----------\nkey: %s\nvalue: %s\n", key, value)
			}
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(valuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// valuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// valuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
