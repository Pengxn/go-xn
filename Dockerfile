# Build stage
FROM golang:alpine AS builder

# ENV GOPROXY https://goproxy.cn,direct

WORKDIR /src

COPY . .

RUN apk add --no-cache build-base make git; \
    go mod download; \
    go mod tidy; \
    CGO_CFLAGS="-D_LARGEFILE64_SOURCE" make

# Server image
FROM alpine:latest

COPY --from=builder /src/build/ /

CMD [ "/go-xn", "web" ]
