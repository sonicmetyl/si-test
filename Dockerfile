FROM golang:1.10-alpine
ENV GOBIN /go/bin

RUN apk --update add git openssh curl
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
COPY ./app /go/src/go-sql/app
WORKDIR /go/src/go-sql/app

COPY ./app/Gopkg.toml ./app/Gopkg.lock ./wait-for-it.sh /go/src/go-sql/app/

RUN dep ensure --vendor-only