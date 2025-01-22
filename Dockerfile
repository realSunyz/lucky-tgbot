FROM golang:1.23.5 AS builder
ARG VERSION=prod
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go build -o lucky

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/lucky .
COPY --from=builder /app/plugin ./plugin

CMD ["./lucky"]
