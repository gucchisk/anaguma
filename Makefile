NAME := anaguma

build: main.go
	go build -v -o bin/$(NAME) github.com/gucchisk/anaguma
