FROM pytorch/pytorch:2.8.0-cuda12.9-cudnn9-runtime

RUN pip install huggingface_hub[cli,hf_xet] transformers grpcio-tools

RUN hf download typeform/distilbert-base-uncased-mnli --include model.safetensors

WORKDIR /run

COPY . .

CMD [ "python3", "/run/classification-service/cmd/main.py" ]
