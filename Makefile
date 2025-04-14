.PHONY: gen_proto

gen_proto:
	protoc --go_out=. --go-grpc_out=. ./protos/collect-service.proto
