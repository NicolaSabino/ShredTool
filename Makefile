.PHONY: build clean

build:
	go build -o bin/shred

install:
	go install 

clean:
	rm -f bin/shred

run:
	go run main.go shred.go

test:
	go test
