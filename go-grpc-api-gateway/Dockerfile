### STAGE 1 ###
FROM golang:1.18-alpine as builder
ENV APP_NAME go-grpc-api-gateway
ENV MAIN_FILE main.go 
ENV MAIN_FILE_PATH cmd 
WORKDIR $GOPATH/src/${APP_NAME}
COPY . .
# RUN go mod tidy
# RUN go mod vendor
RUN go build -o /${APP_NAME} $GOPATH/src/${APP_NAME}/cmd/main.go

### STAGE 2 ###
FROM alpine:3.6
ENV APP_NAME go-grpc-api-gateway
ENV ENVRIONMENT DOCKER
ENV PORT 3000
COPY --from=builder /${APP_NAME} .
EXPOSE ${PORT}
CMD ./${APP_NAME}