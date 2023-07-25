FROM golang:1.20.6

WORKDIR /usr/src/tj-fiber

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy