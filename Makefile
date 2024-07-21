build:
	go build -o ./bin/modular-blockchain-golang

run: build	
	./bin/modular-blockchain-golang

test:
	go test -v ./...

format:
	go fmt ./...	
