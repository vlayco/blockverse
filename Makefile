build:
	@go build -o bin/blockverse

run: build
	@./bin/blockverse

test:
	@go test ./... -v

test_race:
	@go test ./... -v --race

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	 --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

.PHONY: proto