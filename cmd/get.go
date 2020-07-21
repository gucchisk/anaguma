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
	// "github.com/dgraph-io/badger"
	"github.com/spf13/cobra"
	"github.com/gucchisk/anaguma/common"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <key> <dir>",
	Short: "Get value of key",
	Long: "Get value of key",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("require key & badger DB dir")
		}
		var err error
		inputFormat, err = common.NewFormat(in)
		if err != nil {
			return err
		}
		return nil;
	},
	Run: func(cmd *cobra.Command, args []string) {
		key, err := common.StrToByte(args[0], inputFormat)
		if err != nil {
			log.Fatal(err)
		}
		err = db.Open(args[1], verbose)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.Get(key, func(value []byte) error {
			fmt.Printf("%s\n", common.ByteToStr(value, outputFormat))
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.
	getCmd.Flags().StringVarP(&in, "in", "i", "ascii", "input key format")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
