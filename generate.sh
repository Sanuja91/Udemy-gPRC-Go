#!/bin.bash
protoc --go_out=. *.proto

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.