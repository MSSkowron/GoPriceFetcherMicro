build:
	@go build -o bin/pricefetcher

run: build 
	@./bin/pricefetcher

proto:
	protoc --proto_path=proto proto/*.proto --go_out=proto 
	protoc --proto_path=proto proto/*.proto --go-grpc_out=proto 

.PHONY: proto