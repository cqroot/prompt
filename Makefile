.PHONY: test
test:
	go test -v $$(go list ./... | grep -v '/examples/') -covermode=count -coverprofile=coverage.out .

.PHONY: check
check:
	golangci-lint run
	@echo
	gofumpt -l .

.PHONY: cover
cover:
	go tool cover -html=coverage.out

.PHONY: screenshots
screenshots:
	bash $(CURDIR)/scripts/screenshots.bash
