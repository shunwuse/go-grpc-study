```shell
# install gRPG
go get google.golang.org/grpc

# install protoc
# https://github.com/protocolbuffers/protobuf/releases

# install protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# install protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# generate code
protoc --go_out=. --go-grpc_out=. greeter.proto

# run server
run server/*.go

# run client
go run client/main.go
```
