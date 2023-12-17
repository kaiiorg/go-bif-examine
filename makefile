build:
	go build -o ./bin/ ./cmd/go-bif-examine

fmt:
	gofmt -w -s pkg cmd