# ==========================
# Stage 1 - Build
# ==========================
FROM golang:1.26 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o todoapp .

# ==========================
# Stage 2 - Runtime
# ==========================
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/todoapp .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./todoapp"]
