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