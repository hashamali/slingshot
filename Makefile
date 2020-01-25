.PHONY: install run 

install:
	go get

run: install
	go run main.go
