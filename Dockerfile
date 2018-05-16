FROM golang:latest

LABEL maintainer="Ricky Lu <fantasy9830@gmail.com>"

RUN go get github.com/fantasy9830/go-boilerplate

WORKDIR /go/src/github.com/fantasy9830/go-boilerplate

RUN go build

CMD ["./go-boilerplate"]

EXPOSE 8080