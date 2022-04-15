# Blockchain Indexes App

Simple server to interact with Ethereum Blockchain.

## Installation
Install dependencies
```go
go mod download
```

## Binding contract 
```bash
make generate@contract
```

## Usage
### Start server
You can start the server with the following command
```bash
make start
```
or with hot reload
```go
modd
```
or with docker-compose (hot reload available in dev mode)
```bash
make start@docker
```
The server should be available at `localhost:8080`, you'll find the API specs and documentation at `localhost:8080/swagger/index.html`. 

## Tests
Run tests
```bash
make tests
```
## Generate doc
```bash
make generate@doc
```

## Going further 💭
- Add more tests (units/functional)
- Add a CI/CD (e.g: GitHub Actions)
- More .. (Always 😅)