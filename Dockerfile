FROM golang:1.19-alpine

LABEL version="1.0"

LABEL maintainer="Febrilian <febrilian.kr@gmail.com>"

WORKDIR $GOPATH/src/github.com/febriliankr/risetku-shortener-client

COPY . .

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/rest cmd/rest/*.go

EXPOSE 5000

ENTRYPOINT ["/go/bin/rest"]
