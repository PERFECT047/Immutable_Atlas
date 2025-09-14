build: 
	@go build -o bin/ia

run: build
	@./bin/ia

test:
	@go test ./... -v