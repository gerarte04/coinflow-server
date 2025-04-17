.PHONY: gen_proto

gen_proto:
	protoc --go_out=. --go-grpc_out=. ./protos/collection-service.proto
	python3 -m grpc_tools.protoc --python_out=gen/classification_service --grpc_python_out=gen/classification_service -I./protos protos/classification-service.proto
