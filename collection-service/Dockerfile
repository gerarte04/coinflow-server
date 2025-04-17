FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /build

COPY go.sum go.mod ./
RUN go mod download

COPY . .
RUN go build -o service collection-service/cmd/main.go

FROM alpine:3.21.3 AS runner

WORKDIR /app

COPY --from=builder /build/service ./service
COPY --from=builder /build/collection-service/config/config.yaml ./config.yaml

CMD ["/app/service", "--config=/app/config.yaml"]
