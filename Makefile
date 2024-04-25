build:
	@go build -o bin/blockverse

run: build
	@./bin/blockverse

test:
	@go test ./... -v

test_race:
	@go test ./... -v --race