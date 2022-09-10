proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/**/pb/*.proto

server:
	go run cmd/main.go	

check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	go mod vendor && GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger --port=8085 --no-open swagger.yaml
