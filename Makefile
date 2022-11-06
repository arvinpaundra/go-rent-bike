test:
	go test -v -cover ./...

lint: test
	golangci-lint run ./...

build: test lint
	go build -o ./build/main ./internal/app/main.go