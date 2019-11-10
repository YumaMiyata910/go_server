FROM golang:latest

WORKDIR /go/src

COPY server.go .

CMD ["go", "run", "server.go"]