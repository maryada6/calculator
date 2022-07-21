GOPATH = $(go env GOPATH)
build:
	go build -o calci ./cmd/
test :
	mkdir -p coverage/
	go clean -testcache ./... && go test -race ./... -covermode=atomic -coverprofile=coverage/coverage.out
run: build
	./calci