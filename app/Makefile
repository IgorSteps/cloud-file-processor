

## vendor: Download dependencies
vendor: 
	go mod vendor

## lint: Lint code
lint: 
	golangci-lint run

wire:
	wire ./cmd/cloud-file-rest

build:
	go build -o ./dist/cloud-file-rest ./cmd/cloud-file-rest