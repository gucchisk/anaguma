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
	"log"
	"github.com/spf13/cobra"
	"github.com/gucchisk/anaguma/common"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set <key> <value> <dir>",
	Short: "Set value of key",
	Long: "Set value of key",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("require key & value & badger DB dir")
		}
		var err error
		inputFormat, err = common.NewFormat(in)
		if err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		key, err := common.StrToByte(args[0], inputFormat)
		if err != nil {
			log.Fatal(err)
		}
		value, err := common.StrToByte(args[1], inputFormat)
		if err != nil {
			log.Fatal(err)
		}
		err = db.Open(args[2], verbose)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.Set(key, value)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.
	setCmd.Flags().StringVarP(&in, "in", "i", "ascii", "input key & value format")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
