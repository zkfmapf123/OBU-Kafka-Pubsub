# OBU-pubsub

## Execute

```sh
    ## in local
    make up

    ## in local 
    make scp-up
    ## in cloud
    make up
```

## Issue

- Golang Kafka Client 이슈

    - 이게문제가 뭐냐면
    - go-kafka 같은 경우에는 C 라이브러리를 사용
    - Go자체가 C 라이브러리를 사용하다록 설정이 되어있어야 함
    - 그거와 더불어 몇몇개의 라이브러리를 사용해야 함 (pkgconf...)
    - <a href="https://github.com/confluentinc/confluent-kafka-go#getting-started">Go-Kafka</a>

```Dockerfile
## Build Stage
FROM golang:1.21.5-alpine as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk --no-cache update && \
    apk --no-cache add git gcc libc-dev librdkafka-dev vips-dev pkgconf

## CGO_ENABLED : C 라이브러리 사용
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    CGO_ENABLED=1
# GOARCH=amd64

# RUN go build -o main .
# RUN go build -tags musl -o main .
RUN go build -tags dynamic -o main .

## Runner Stage
FROM scratch

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/main .

CMD ["./main"]
```