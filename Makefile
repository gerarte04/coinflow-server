.PHONY: gen_proto

py_grpc_path = gen/classification_service/python

gen_proto:
	protoc --go_out=. --go-grpc_out=. ./protos/collection-service.proto
	protoc --go_out=. --go-grpc_out=. ./protos/classification-service.proto
	protoc --go_out=. --go-grpc_out=. ./protos/storage-service.proto
	protoc --go_out=. --go-grpc_out=. ./protos/auth-service.proto
	python3 -m grpc_tools.protoc --pyi_out=$(py_grpc_path) --python_out=$(py_grpc_path) --grpc_python_out=$(py_grpc_path) -I./protos protos/classification-service.proto

build_services:
	docker compose build

launch_services: build_services
	docker compose up --force-recreate

launch_services_with_tests: build_services
	docker compose --profile test up --force-recreate --abort-on-container-exit --exit-code-from tester

stop_services:
	docker compose down -v
