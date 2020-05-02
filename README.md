# Golang Gateway Microservice Template

Exemplary Golang implementation of a Gateway Microservice.

## Overview

The microservice gateway proxies incoming HTTP requests to subsequent microservices. It is the single entry point for the backend.
The gateway can be extended with middlewares (for authentication, authorisation, etc.)

- Reverse proxy to internal microservices
- Flexible configuration
- Echo HTTP router
- Middlewares for proxy functionality and more
- Mock creation of interfaces with Mockery
- API documentation with Swagger UI and go-swagger
- Code linting with GolangCI

## Run the microservice

To run the server, follow these simple steps:

```bash
go run main.go
```

To run the server in a docker container

```bash
docker build -t gateway-service .
```

Once image is built use

```bash
docker run --rm -it gateway-service
```

## Swagger specification

1. Download the latest release of [go-swagger](https://github.com/go-swagger/go-swagger/releases)

2. Generate the Swagger specification with

    ```bash
    %GOPATH%/bin/go-swagger generate spec --scan-models -o ./swaggerui/swagger.json
    ```

3. The interactive Swagger documentation is available via the [/swagger endpoint](http://localhost:8080/swagger/).

## Environment variable

API_VERSION: the version parameter in the endpoint urls

Launch.json contains a debug launch configuration. Environment parameters can be specified in a local.env file.


## Linting

Install the linter runner [GolangCI-Lint](https://github.com/golangci/golangci-lint). The config file is already part of this repository.

## Unit tests and mocks

[Mockery](https://github.com/vektra/mockery) is used to generate mocks from interfaces for unit tests.

```bash
%GOPATH%/bin/mockery -all -case=underscore -inpkg
```

Run all tests:

```bash
go test ./...
```