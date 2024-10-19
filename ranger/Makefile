#!/usr/bin/make -f

test:
	go fmt ./...
	go mod tidy
	go test -cover -timeout=1s -race -count 10 ./...

onefile:
	@go-mergepkg -dirs "." -header "github.com/mdw-go/funcy/ranger@$(shell git describe) (a little copy-paste is better than a little dependency)"
