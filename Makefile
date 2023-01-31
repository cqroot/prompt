.PHONY: check
check:
	golangci-lint run
	@echo
	gofumpt -l .
