package main

//go:generate protoc -I. --go_out=./schema  --go_opt=paths=source_relative core_search.proto
//go:generate protoc -I. --go-grpc_out=./schema  --go-grpc_opt=paths=source_relative  core_service.proto
