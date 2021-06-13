.PHONY: run build-linux build-darwin

run:
	go run ./cmd/main/main.go

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o _output/catchall ./cmd/main/main.go

build-darwin:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o _output/catchall ./cmd/main/main.go
