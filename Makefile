.PHONY: start build

all: run

build:
	swag init
	GOOS=linux GOARCH=amd64 go build ./main.go && mv main start
run: 
	go run main.go
clean:
	go clean
swag:
	swag init