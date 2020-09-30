# anaguma

anaguma is command line interface for [BadgerDB](https://github.com/dgraph-io/badger)


### Install

```
$ brew tap gucchisk/tap
$ brew install anaguma
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
  -b, --bversion uint8   badger version [1, 2] (default 1)
  -h, --help             help for anaguma
  -l, --log              log output
  -o, --out string       output format [ascii hex base64] (default "ascii")
  -v, --version          show anaguma version

Use "anaguma [command] --help" for more information about a command.
```
