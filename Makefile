.PHONY: build clean

build:
	go build -o bin/shred

clean:
	rm -f bin/shred

test:
	go test
