.PHONY: install test binary cover cover-html cover-total

binary:
	@go build ./cmd/backend

install:
	@go install ./...

test:
	@go test ./... -cover

cover:
	@go test -coverprofile=cover.out ./...

cover-html:
	@go test -coverprofile=cover.out ./...
	@go tool cover -html=cover.out

cover-total:
	@go test -coverprofile=cover.out ./...
	@go tool cover -func cover.out | grep total:

lint:
	golangci-lint run ./...
