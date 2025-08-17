FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN  apk add --no-cache git

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest AS runner

WORKDIR /root/

COPY --from=builder /app/main ./main
COPY --from=builder /app/migrations ./migrations

EXPOSE 9000

CMD ["./main"]