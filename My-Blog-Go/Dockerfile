# ========== Stage 1: Build Go Binary ==========
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server .

# ========== Stage 2: Run ==========
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/static ./static

EXPOSE 8000
CMD ["./server"]
