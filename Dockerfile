FROM golang:1.15-alpine

WORKDIR /go/pMade.api

COPY src/app ./src/app
COPY go.mod .

RUN apk add --no-cache git \
  && go get github.com/oxequa/realize

WORKDIR /go/pMade.api/src/app

RUN go build -o app