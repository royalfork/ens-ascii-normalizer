.PHONY: all
all: generate

.PHONY: generate
generate:
	@abigen --sol ./contracts/AsciiNormalizer.sol --pkg bindings --out ./bindings/AsciiNormalizer.go

.PHONY: test
test: generate
	@go test ./test $(TEST_ARGS)
