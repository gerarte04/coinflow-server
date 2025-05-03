FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /build

COPY go.sum go.mod ./
RUN go mod download

COPY . .
RUN go build -o api api-gateway/cmd/main.go

FROM alpine:3.21.3 AS runner

WORKDIR /app

RUN apk add --no-cache curl

COPY --from=builder /build/api ./api
COPY --from=builder /build/api-gateway/config/config.yaml ./config.yaml

CMD ["/app/api", "--config=/app/config.yaml"]
