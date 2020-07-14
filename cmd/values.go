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
	"github.com/dgraph-io/badger"
	"github.com/spf13/cobra"
)

// valuesCmd represents the values command
var valuesCmd = &cobra.Command{
	Use:   "values",
	Short: "Show keys & values",
	Long: "Show keys & values",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("require badger DB dir")
		}
		return nil;
	},
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		fmt.Printf("badger DB dir: %s\n", dir)
		db, err := badger.Open(badger.DefaultOptions(dir))
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		err = db.View(func(txn *badger.Txn) error {
			opts := badger.DefaultIteratorOptions
			opts.PrefetchSize = 10
			it := txn.NewIterator(opts)
			defer it.Close()
			for it.Rewind(); it.Valid(); it.Next() {
				item := it.Item()
				k := item.Key()
				err = item.Value(func(v []byte) error {
					fmt.Printf("%s: %v\n", k, v)
					return nil
				})
				if err != nil {
					return err
				}
			}
			return nil
		})
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
