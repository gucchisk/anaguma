NAME := anaguma
BIN := bin

build:
	go build -v -o $(BIN)/$(NAME) github.com/gucchisk/anaguma

v%: 
	go run scripts/release/main.go $@

pack:
	go run scripts/package/main.go

clean:
	rm -rf $(BIN)

.PHONY: clean
