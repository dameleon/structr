VERSION = $$(git describe --tags --always --dirty) ($$(git name-rev --name-only HEAD))

BUILD_FLAGS = -ldflags "-X main.Version \"$(VERSION)\" "

deps:
	go get -d
destdeps:
	go get -d -t
test: testdeps
	go test ./...
build: deps
	go build $(BUILD_FLAGS)
install: deps
	go install $(BUILD_FLAGS)

.PHONY: deps build install
