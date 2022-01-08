FROM golang:1.17-alpine

WORKDIR /app

COPY . .

CMD CGO_ENABLED=0 go test -v ./...
