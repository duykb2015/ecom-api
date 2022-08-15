FROM golang:1.18-alpine3.16 as modules
COPY go.mod go.sum /modules/

WORKDIR /modules
RUN go mod download
