.PHONY: build test clean

build:
	go build -o ascii-art ./cmd/ascii-art

test:
	go test ./...

clean:
	rm -f ascii-art ascii-art.exe coverage.out

install:
	go install ./cmd/ascii-art

lint:
	gofmt -s -l .
	go vet ./...