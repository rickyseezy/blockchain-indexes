# Base image
FROM golang:alpine3.13 as base

ENV GOPATH /go
ENV GO111MODULE on
ENV GOSUMDB off

RUN apk add build-base

WORKDIR /go/src/github.com/rickyseezy/block

COPY . .

# modd
RUN go get github.com/cortesi/modd/cmd/modd

RUN go mod download

ENTRYPOINT ["modd"]

EXPOSE 8080

# builder
FROM base as builder

RUN go build -v -o app cmd/blockindex/main.go

# runtime
FROM alpine:3.13.6 as runtime

# Switch working directory to /usr/bin
WORKDIR /usr/bin

# Copies the binary file from the BUILD container to /usr/bin
COPY --from=builder /go/src/github.com/rickyseezy/block/app .
COPY --from=builder /go/src/github.com/rickyseezy/block/internal/config/production.env .

EXPOSE 8080

ENTRYPOINT ["/usr/bin/app"]