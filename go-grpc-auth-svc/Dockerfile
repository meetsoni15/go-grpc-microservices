### PHASE 1 ###
FROM golang:1.18-alpine as builder

ENV APP_NAME=go-grpc-auth-svc MAIN_FILE=main.go MAIN_FILE_PATH=cmd

WORKDIR $GOPATH/src/${APP_NAME}

COPY . . 

RUN go build -v -o /${APP_NAME} $GOPATH/src/${APP_NAME}/${MAIN_FILE_PATH}/${MAIN_FILE}

### PHASE 2 ###
FROM alpine:3.6

ENV APP_NAME=go-grpc-auth-svc PORT=50051

COPY --from=builder /${APP_NAME} .

EXPOSE ${PORT}

CMD ./${APP_NAME}