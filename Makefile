.PHONY: format-check lint test test-integration verify install-hooks

format-check:
	test -z "$$(gofmt -l .)"

lint:
	golangci-lint run

test:
	go test ./... -coverprofile=coverage.out

test-integration:
	go test ./... -tags=integration -count=1

verify: format-check lint test

install-hooks:
	git config core.hooksPath .githooks
