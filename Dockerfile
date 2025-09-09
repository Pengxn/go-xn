# Build stage
FROM golang:1.25-alpine AS builder

# ENV GOPROXY https://goproxy.cn,direct

WORKDIR /src

COPY . .

RUN apk add --no-cache build-base make git; \
    go mod download; \
    go mod tidy; \
    make

# Server image
FROM alpine:latest

COPY --from=builder /src/build/ /

CMD [ "/go-xn", "web" ]
