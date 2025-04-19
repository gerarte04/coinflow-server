FROM python:3.13.3-slim-bookworm

RUN pip install torch --index-url https://download.pytorch.org/whl/cpu
RUN pip install hf_xet transformers grpcio-tools

COPY . .

CMD [ "python3", "classification-service/cmd/main.py" ]
