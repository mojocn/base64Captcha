.PHONY: test

test:
	go test $(shell go list ./... | grep -Ev '/vendor/|/test/' )