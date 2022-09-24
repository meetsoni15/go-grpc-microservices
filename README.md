# Go GRPC Microservice
GO Grpc Microservice

## Description
These demo project contains GRPC Microservices and API gateway with docker and k8s implementations.
All of mentioned microservices have their Dockerfile inside directory

#### Here are different project directory for microservices 

- https://github.com/meetsoni15/go-grpc-project/tree/master/go-grpc-auth-svc - Authentication SVC (gRPC)
- https://github.com/meetsoni15/go-grpc-project/tree/master/go-grpc-api-gateway - API Gateway (HTTP)
- https://github.com/meetsoni15/go-grpc-project/tree/master/go-grpc-product-svc - Product SVC (gRPC)
- https://github.com/meetsoni15/go-grpc-project/tree/master/go-grpc-order-svc - Order SVC (gRPC)

> utils folder contaning docker compose and k8s configuration files


## Running the app inside docker container

```bash
$ cd utils
$ docker-compose up --build
```

## Running the app inside k8s cluster

```bash
$ cd utils
$ kubectl -f k8s/
```

> Here I've used local images of different microservices rather than from DockerHub
