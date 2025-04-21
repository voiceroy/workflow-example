FROM golang:1.24.2

WORKDIR /usr/src/app

COPY main.go main.go
COPY main_test.go main_test.go
COPY go.mod go.mod

RUN go build -o /usr/local/bin/main

ENTRYPOINT ["main"]
