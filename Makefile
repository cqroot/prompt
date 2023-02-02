.PHONY: test
test:
	go test -v ./...

.PHONY: check
check:
	golangci-lint run
	@echo
	gofumpt -l .
