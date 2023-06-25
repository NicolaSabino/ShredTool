.PHONY: build clean

build:
	go build -o bin/shred

clean:
	rm -f bin/shred

install:
	go install 

run:
	go run main.go shred.go

test:
	go test -v -cover

cov_report:
	go test -v -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out
