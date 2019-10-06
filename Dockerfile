FROM golang:1.13-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get -u github.com/golang/dep/...

COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

EXPOSE 5000

