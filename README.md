# anaguma

anaguma is command line interface for [BadgerDB](https://github.com/dgraph-io/badger)


### Install

```
$ go get github.com/gucchisk/anaguma
```

### Usage

```
anaguma is a CLI tool to access badger DB

Usage:
  anaguma [command]

Available Commands:
  get         Get value of key
  help        Help about any command
  keys        Show keys
  set         Set value of key
  values      Show keys & values

Flags:
  -h, --help            help for anaguma
  -l, --log             log output
  -o, --out string      output format [ascii hex base64] (default "ascii")
  -v, --version uint8   badger version [1, 2] (default 1)

Use "anaguma [command] --help" for more information about a command.
```
