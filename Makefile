test: acceptance

acceptance:
	@echo "> Running acceptance tests..."
	go test -v -count=1 ./acceptance/...

.PHONY: acceptance
