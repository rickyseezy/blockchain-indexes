# bind contract
generate@contract:
	abigen --abi pkg/abi/contract.abi --pkg abi --type Contract --out pkg/abi/contract.go

# doc
generate@doc:
	swag init -g ./cmd/blockindex/main.go --parseDependency --parseInternal --parseDepth 1

# server
start:
	go run cmd/blockindex/main.go

start@docker:
	docker-compose up app

stop:
	docker-compose down

# tests
tests:
	go test ./...
