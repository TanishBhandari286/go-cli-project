FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o address-book cmd/main.go

FROM alpine:latest

RUN apk add --no-cache procps

WORKDIR /app
COPY --from=builder /app/address-book .
COPY healthcheck.sh .
RUN chmod +x healthcheck.sh

RUN mkdir -p data
VOLUME /app/data

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD /app/healthcheck.sh

ENTRYPOINT ["./address-book"] 