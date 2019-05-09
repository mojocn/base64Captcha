.PHONY: test
PACKAGES = $(shell find ./ -type d -not -path '*/\.*')

test:
	go test $(shell go list ./... | grep -Ev '/vendor/|/test/' )

travis:
	echo "mode: atomic" > coverage.txt
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=atomic $(pkg);\
		tail -n +2 coverage.out >> coverage.txt;)
