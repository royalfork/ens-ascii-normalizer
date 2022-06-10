.PHONY: all
all: generate

.PHONY: generate
generate:
	@abigen --sol ./contracts/ENSAsciiNormalizer.sol --pkg bindings --out ./bindings/ENSAsciiNormalizer.go

.PHONY: test
test: generate
	@go test ./test $(TEST_ARGS)
