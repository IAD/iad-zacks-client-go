lint:
	docker run --rm -v $(PWD):$(PWD) -w $(PWD) -u `id -u $(USER)` -e GOLANGCI_LINT_CACHE=/tmp/.cache -e GOCACHE=/tmp/.cache golangci/golangci-lint:v1.53.3 golangci-lint run -v --fix

test:
	go test ./...
