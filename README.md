### Notes
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

```shell
# on ubuntu
sudo apt install openssl

# generate private key
openssl genrsa -out server.key 2048

# generate certificate signing request
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
openssl req -new -key server.key -out server.csr

# openssl.cnf
cp /etc/ssl/openssl.cnf openssl.cnf
# uncomment "copy_extensions = copy"
# uncomment "req_extensions = v3_req"
# add "subjectAltName = @alt_names" under [ v3_req ]
# add [ alt_names ]
# add "DNS.1 = localhost" under [ alt_names ]

# generate certificate
openssl genpkey -algorithm RSA -out test.key -pkeyopt rsa_keygen_bits:2048
openssl req -new -nodes -key test.key -out test.csr -days 3650 -subj "/C=US/ST=CA/L=San Francisco/O=myComp/OU=myOrg/CN=myName" -config openssl.cnf -extensions 'v3_req'
openssl x509 -req -in test.csr -CA server.crt -CAkey server.key -CAcreateserial -out test.crt -days 3650 -sha256 -extfile openssl.cnf -extensions 'v3_req'

```

```shell
# generate cert
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```