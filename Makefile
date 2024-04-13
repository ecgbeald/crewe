run: build
	@./bin/api

build:
	@go build -o bin/api

# ... means go recursively
# test:
# 	@go test -v ./...