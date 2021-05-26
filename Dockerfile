FROM golang:alpine AS builder

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /src

COPY . .

RUN apk add --no-cache build-base make git; \
    go mod download; \
    cp src/config/example.ini $HOME/fyj.ini; \
    make && make web && make test

FROM alpine:latest

WORKDIR /fyj

COPY --from=builder /src/build/ /fyj/

CMD [ "./go-xn", "web" ]
