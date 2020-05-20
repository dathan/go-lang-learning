# Go parameters

.PHONY: all
all: lint test

.PHONY: lint
lint:
				golangci-lint run ./...


.PHONY: test
test:
				go test -p 6 -v -covermode=count -coverprofile=./coverage.out ./...

.PHONY: clean
clean:
				go clean
				find . -type d -name '.tmp_*' -prune -exec rm -rvf {} \;

.PHONY: vendor
vendor:
				go mod vendor