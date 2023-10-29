build:
	@go build -o bin/gobank

run: build
	@./bin/gobank

test: 
	@go test -v ./... #test all files in the root directory
	