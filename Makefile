.PHONY: gen_proto

py_grpc_path = gen/classification_service/python

gen_proto:
	protoc --go_out=. --go-grpc_out=. ./protos/collection-service.proto
	protoc --go_out=. --go-grpc_out=. ./protos/classification-service.proto
	python3 -m grpc_tools.protoc --pyi_out=$(py_grpc_path) --python_out=$(py_grpc_path) --grpc_python_out=$(py_grpc_path) -I./protos protos/classification-service.proto
