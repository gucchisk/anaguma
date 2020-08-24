NAME := anaguma
BIN := bin

build:
	go build -v -o $(BIN)/$(NAME) github.com/gucchisk/anaguma

v%: 
	go run script/release.go $@

clean:
	rm -rf $(BIN)

.PHONY: clean
