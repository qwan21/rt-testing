BINARY_NAME=rt
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BINARY_NAME)_windows.exe

all: test build
build:
	@go build -o $(BINARY_NAME) ./cmd/server
test:
	@go test -v -cover ./...
clean:
	@go clean
	@rm -rf $(BINARY_NAME)
	@rm -rf $(BINARY_UNIX)
	@rm -rf $(BINARY_WINDOWS)
clean-all: clean
	@rm -rf rt*.tar
run:
	@go build -o $(BINARY_NAME)
	./$(BINARY_NAME)
get-artifact: cross-build
	@tar -cvf $(BINARY_NAME)_windows.tar $(BINARY_WINDOWS) config/config.yml
	@tar -cvf $(BINARY_NAME)_unix.tar $(BINARY_UNIX) config/config.yml
	@make clean


# cross compilation
cross-build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_UNIX) -v ./cmd/rt
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_WINDOWS) -v ./cmd/rt
# docker-build:
# 	docker run --rm -it -v "$(GOPATH)":/go -w ~//src/rt golang:latest go build -o "$(BINARY_UNIX)" -v ./cmd/rt