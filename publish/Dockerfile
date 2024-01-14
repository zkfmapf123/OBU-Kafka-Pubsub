## Build Stage
FROM golang:1.21.5-alpine as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk --no-cache update && \
    apk --no-cache add git gcc libc-dev librdkafka-dev vips-dev pkgconf

RUN go build -tags dynamic -o main .

## Runner Stage
FROM alpine:latest

WORKDIR /usr/src/app

RUN apk --no-cache add libc-dev librdkafka-dev pkgconf
COPY --from=builder /usr/src/app/main .

CMD ["./main"]