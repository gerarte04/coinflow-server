import sys
sys.path.append(".")
sys.path.append("./classification-service")
sys.path.append("./gen/classification_service/python")

from concurrent import futures
from gen.classification_service.python import classification_service_pb2_grpc
import grpc
import internal.api.servicer as servicer
import internal.usecases.service as service
import os
import torch

def main():
	host = os.getenv("GRPC_CLASSIFICATION_SERVICE_HOST")
	port = os.getenv("GRPC_CLASSIFICATION_SERVICE_PORT")

	device = torch.device('cuda' if torch.cuda.is_available() else 'cpu')
	cat_service = service.CategoryService(os.getenv("CLASSIFICATION_MODEL_NAME"), device)
	svcr = servicer.ClassificationServicer(cat_service)

	server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
	classification_service_pb2_grpc.add_ClassificationServicer_to_server(svcr, server)

	server.add_insecure_port(f'{host}:{port}')
	server.start()
	print(f'server started at {host}:{port}', flush=True)
	server.wait_for_termination()

if __name__ == "__main__":
	main()
