#!/bin.bash

# Converts all .proto files into Go
protoc --go_out=. *.proto

# Converts specified proto file into gRPC in Go
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.


# Runs Go Server
go run greet/greet_server/server.go

# Runs Go Client
go run greet/greet_client/client.go
