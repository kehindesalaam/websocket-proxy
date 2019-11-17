FROM golang:1.12

ENV APP_DIR /go/src/github.com/terawork-com/message-socket-service

RUN go get github.com/githubnemo/CompileDaemon

COPY . ${APP_DIR}
WORKDIR ${APP_DIR}
