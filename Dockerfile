# -----------------------------
# Stage 1 -> Build Stage
# -----------------------------
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

# -----------------------------
# Stage 2 -> Production Stage
# -----------------------------
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./main"]