# Build stage
FROM golang:alpine AS builder

# ENV GOPROXY https://goproxy.cn,direct

WORKDIR /src

COPY . .

RUN apk add --no-cache build-base make git; \
    go mod download; \
    cp src/config/example.ini $HOME/fyj.ini; \
    make && make web

# Server image
FROM alpine:latest

COPY --from=builder /src/build/ /fyj/

CMD [ "/fyj/go-xn", "web" ]
