FROM golang:1.14 AS build

WORKDIR /go/src/gateway

ENV GOPATH /go
ENV CGO_ENABLED 0
ENV GO111MODULE on

COPY . ./gateway-service

RUN cd gateway-service && \
    go vet ./... && \
    go test -v ./... -coverprofile=cover.out && go tool cover -func=cover.out && \
    go build -a -installsuffix cgo -o app .

FROM alpine:3.10 AS runtime
COPY --from=build /go/src/gateway/gateway-service/app ./
COPY --from=build /go/src/gateway/gateway-service/swaggerui ./swaggerui
EXPOSE 8080/tcp
ENTRYPOINT ["./app"]
