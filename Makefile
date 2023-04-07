.DEFAULT_GOAL:=help

## === Tasks ===

.PHONY: gomod
## Run go mod tidy
gomod:
	go mod tidy

.PHONY: outdated
## Show outdated direct dependencies
outdated:
	go list -u -m -json all | go-mod-outdated -update -direct

.PHONY: check
## Run tests and linters
check: test lint

.PHONY: test
## Run tests
test:
	gotestsum --format dots-v2 ./...

.PHONY: bench
## Run benchmarks
bench:
	go test -run ^$$ -bench . | prettybench

.PHONY: lint
## Run linter
lint:
	golangci-lint run

.PHONY: lint.fix
## Fix lint violations
lint.fix:
	golangci-lint run --fix

## === Utils ===

## Show help text
help:
	@awk '{ \
			if ($$0 ~ /^.PHONY: [a-zA-Z\-\_0-9]+$$/) { \
				helpCommand = substr($$0, index($$0, ":") + 2); \
				if (helpMessage) { \
					printf "\033[36m%-23s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^[a-zA-Z\-\_0-9.]+:/) { \
				helpCommand = substr($$0, 0, index($$0, ":")); \
				if (helpMessage) { \
					printf "\033[36m%-23s\033[0m %s\n", \
						helpCommand, helpMessage"\n"; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^##/) { \
				if (helpMessage) { \
					helpMessage = helpMessage"\n                        "substr($$0, 3); \
				} else { \
					helpMessage = substr($$0, 3); \
				} \
			} else { \
				if (helpMessage) { \
					print "\n                        "helpMessage"\n" \
				} \
				helpMessage = ""; \
			} \
		}' \
		$(MAKEFILE_LIST)
