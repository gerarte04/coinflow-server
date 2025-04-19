import sys
sys.path.append("./classification-service/internal")

from gen.classification_service import classification_service_pb2_grpc
from gen.classification_service.classification_service_pb2 import GetTextCategoryRequest, GetTextCategoryResponse
import grpc
from usecases.object import CategoryServiceInterface

class ClassificationServicer(classification_service_pb2_grpc.ClassificationServicer):
	def __init__(self, service: CategoryServiceInterface):
		super().__init__()
		self.service = service

	def GetTextCategory(self, request, context):
		try:
			category = self.service.DetectCategory(
				request.GetTextCategoryRequest.text,
				request.GetTextCategoryRequest.labels
			)

			return GetTextCategoryResponse(
				category = category
			)
		except:
			_, value, _ = sys.exc_info()
			context.set_code(grpc.StatusCode.INTERNAL)
			context.set_details(value.strerror)
