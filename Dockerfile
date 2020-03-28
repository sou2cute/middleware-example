ARG GO_VERSION=1.13.9
ARG PORT=8080

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && \
    apk --no-cache add git alpine-sdk

WORKDIR /middleware-example

RUN go mod init github.com/sou2cute/middleware-example
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/go-redis/redis/v7
RUN go get github.com/go-redis/redis_rate/v8
COPY main.go .

RUN go build -o ./main ./main.go

FROM alpine:latest

ENV PORT=${PORT}

RUN apk update && \
    apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /middleware-example/main .

EXPOSE ${PORT}

ENTRYPOINT [ "./main" ]