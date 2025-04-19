import sys
sys.path.append("./classification-service/internal")

from transformers import pipeline
from usecases.object import CategoryServiceInterface

class CategoryService(CategoryServiceInterface):
	def __init__(self):
		super().__init__()
		self.clf = pipeline("zero-shot-classification")

	def DetectCategory(self, text, labels):
		if len(labels) == 0:
			return ""

		pred = self.clf(text, labels)
		
		max_value = pred['scores'][0]
		argmax_category = pred['labels'][0]

		for i in range(1, len(pred['scores'])):
			if pred['scores'][i] > max_value:
				max_value = pred['scores'][i]
				argmax_category = pred['labels'][i]

		return argmax_category
