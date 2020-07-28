FROM golang:latest

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app

RUN go get -v
RUN go mod download

#RUN go get "github.com/githubnemo/CompileDaemon"

#ENTRYPOINT CompileDaemon -directory="/go/bin/app"

EXPOSE ${APP_PORT}