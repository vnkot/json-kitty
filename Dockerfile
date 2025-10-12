FROM golang:1.24.3-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /server ./
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

EXPOSE 8000
CMD ["./server"]