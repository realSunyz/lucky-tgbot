FROM golang:1.23.5-bookworm AS builder
ARG VERSION=prod
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o lucky

FROM debian:bookworm
WORKDIR /app
COPY --from=builder /app/lucky .
COPY --from=builder /app/plugin ./plugin

CMD ["./lucky"]
