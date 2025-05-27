.PHONY: run build test clean

BINARY_NAME=go-route-planner

build:
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

test:
	go test ./...

clean:
	go clean
