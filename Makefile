
fmt:
	go fmt ./cmd/web

run: 
	go run ./cmd/web

watch:
	air .

build: fmt
	go build ./cmd/web

revise:
	@goimports-reviser -rm-unused -set-alias -format -recursive .


