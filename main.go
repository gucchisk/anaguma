/*

anaguma is a CLI tool to access badger DB

Usage

usage:

  anaguma [command] [flag]

Flags for all commands

flags:

  -b, --bversion uint8   badger version [1, 2] (default 1)
  -h, --help             help for anaguma
  -l, --log              log output
  -o, --out string       output format [ascii hex base64] (default "ascii")
  -v, --version          show anaguma version

Command

get:

  anaguma get <key> <dir> [flag]

  -i, --in string   input key format (default "ascii")

help:

  anaguma help [command] [flag]

keys:

  anaguma keys <dir> [flag]

set:

  anaguma set <key> <value> <dir> [flags]

  -i, --in string   input key & value format (default "ascii")

values:

  anaguma values <dir> [flags]

*/
package main

import "github.com/gucchisk/anaguma/cmd"

func main() {
	cmd.Execute()
}
