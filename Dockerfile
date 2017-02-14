FROM golang:1.7
MAINTAINER Hussein Galal hussein@rancher.com

EXPOSE 8000

ADD . /go/src/simpleInterceptor
WORKDIR /go/src/simpleInterceptor

RUN go get
RUN go install

ENTRYPOINT ["/go/bin/simpleInterceptor"]
