FROM python:3.13.3-slim-bookworm

RUN pip install torch --index-url https://download.pytorch.org/whl/cpu
RUN pip install huggingface_hub[cli,hf_xet] transformers grpcio-tools

RUN huggingface-cli download typeform/distilbert-base-uncased-mnli --include model.safetensors

WORKDIR /run

COPY . .

CMD [ "python3", "/run/classification-service/cmd/main.py" ]
