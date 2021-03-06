VERSION = $$(git describe --tags --always --dirty) ($$(git name-rev --name-only HEAD))

BUILD_FLAGS = -ldflags "-s -w -X \"main.Version=$(VERSION)\" "

resources:
	go get -u github.com/jteeuwen/go-bindata/...
	go-bindata -o resources.go resources/
deps: resources
	go get -d ./...
testdeps: resources
	go get -d -t ./...
test: testdeps
	go test ./...
build: deps
	go build $(BUILD_FLAGS)
install: deps
	go install $(BUILD_FLAGS)

.PHONY: resources deps testdeps test build install
