
fmt:
	go fmt ./cmd/web

run: 
	go run ./cmd/web

build: fmt
	go build ./cmd/web


