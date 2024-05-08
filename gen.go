package main

//go:generate protoc -I. --go_out=.  --go_opt=paths=source_relative core_search.proto
//go:generate protoc -I. --go-grpc_out=.  --go-grpc_opt=paths=source_relative  core_service.proto
