# Build stage
FROM golang:alpine AS builder

# ENV GOPROXY https://goproxy.cn,direct

WORKDIR /src

COPY . .

RUN apk add --no-cache build-base make git; \
    go mod download; \
    make && make web

# Server image
FROM alpine:latest

COPY --from=builder /src/build/ /

CMD [ "/go-xn", "web" ]
