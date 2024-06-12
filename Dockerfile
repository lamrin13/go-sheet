FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/api-server/main.go
CMD ["./main"]