import sys
sys.path.append("./classification-service/internal")

from gen.classification_service.python import classification_service_pb2_grpc
from gen.classification_service.python.classification_service_pb2 import GetTextCategoryRequest, GetTextCategoryResponse
import grpc
from usecases.object import CategoryServiceInterface

class ClassificationServicer(classification_service_pb2_grpc.ClassificationServicer):
	def __init__(self, service: CategoryServiceInterface):
		super().__init__()
		self.service = service

	def GetTextCategory(self, request, context):
		try:
			category = self.service.DetectCategory(
				request.text,
				request.labels
			)

			return GetTextCategoryResponse(
				category = category
			)
		except Exception as e:
			context.set_code(grpc.StatusCode.INTERNAL)
			context.set_details(repr(e))
