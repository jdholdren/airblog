.PHONY: install test build serve clean pack deploy ship

test: install
	go test ./... -coverprofile=coverage.out
