.PHONY: gen_proto

py_grpc_path = gen/classification_service/python
desc_set_path = gen/descriptor_set.pb

gen_proto:
	-git clone https://github.com/googleapis/googleapis
	protoc -Igoogleapis -I./protos --go_out=. --go-grpc_out=. collection-service.proto classification-service.proto storage-service.proto auth-service.proto
	python3 -m grpc_tools.protoc --pyi_out=$(py_grpc_path) --python_out=$(py_grpc_path) --grpc_python_out=$(py_grpc_path) -I./protos protos/classification-service.proto

build_services:
	docker compose build

launch_services: build_services
	docker compose up --force-recreate

launch_services_with_tests: build_services
	docker compose --profile test up --force-recreate --abort-on-container-exit --exit-code-from tester

stop_services:
	docker compose down -v
