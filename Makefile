.PHONY: build run test clean

build:
	go build .

run:
	go run main.go

test:
	go test -v ./...

clean:
	rm -f hackernews
