test:
	go test -v ./...

build:
	go build -o ./tmp/main cmd/cli/main.go

run:
	go run cmd/cli/main.go

