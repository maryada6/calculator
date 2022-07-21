GOPATH = $(go env GOPATH)
build:
	go build -o calci ./cmd/
test :
	go test ./....
run: build
	./calci