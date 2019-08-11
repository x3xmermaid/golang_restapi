FROM golang:alpine

MAINTAINER try satria try.satria.a@gmail.com

ADD . /go/src/golang_restapi

ARG appname=golang_restapi
ARG http_port=8080

RUN apk update && \
    apk upgrade && \
    apk add git
	
RUN go get -d github.com/gorilla/mux
RUN go get -d github.com/go-sql-driver/mysql


RUN mkdir -p /go/src/golang_restapi
COPY . /go/src/golang_restapi
WORKDIR /go/src/golang_restapi

RUN go build main.go
CMD ["./main"]

EXPOSE $http_port