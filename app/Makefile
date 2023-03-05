setup:
	go mod download
	go mod tidy

build-tools: setup
	export CGO_ENABLED=0
	go build -o ./bin/gotestsum gotest.tools/gotestsum

test: build-tools
	bin/gotestsum ./...

ci-test:
	go test -race -p 1 ./... -coverprofile=coverage.out
