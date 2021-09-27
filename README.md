# oapi-sample
OpenAPI and sqlc sample project

## Requirements

- Go 1.17
- oapi-codegen
- sqlc
- postgresql 13

## Setup

```shell
cd .devcontainer
docker compose up -d

go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.8.2
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```
## Build

```shell
go build cmd/server/main.go
```

## Run

```shell
go run cmd/server/main.go
```

## Generate

Updates docs/openapi.yaml or sql//**.sql. generate.

```shell
go generate
```

or

for docs/openapi.yaml

```shell
oapi-codegen --config=types.cfg.yaml docs/openapi.yaml
oapi-codegen --config=server.cfg.yaml docs/openapi.yaml
```

for sql//**.sql

```shell
sqlc generate -f sqlc.yaml
```