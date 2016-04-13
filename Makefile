VERSION = $$(git describe --tags --always --dirty) ($$(git name-rev --name-only HEAD))

BUILD_FLAGS = -ldflags "-X main.Version \"$(VERSION)\" "

deps:
	go get -d
testdeps:
	go get -d -t
test: testdeps resources
	go test ./...
resources:
	go-bindata -o resources.go resources/
build: deps resources
	go build $(BUILD_FLAGS)
install: deps resources
	go install $(BUILD_FLAGS)

.PHONY: resources deps testdeps test build install
