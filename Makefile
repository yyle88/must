COVERAGE_DIR ?= .coverage

# cp from: https://github.com/yyle88/done/blob/d7658fd6478b77a7c4dfc50b1cea36134ad90b02/Makefile#L4
test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
