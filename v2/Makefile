#!/usr/bin/make -f

test:
	go fmt ./...
	go mod tidy
	go test -cover -timeout=1s -race -count 1 ./...

onefile:
	@go-mergepkg -dirs "." -header "github.com/mdwhatcott/funcy/v2@$(shell git describe) (a little copy-paste is better than a little dependency)"
