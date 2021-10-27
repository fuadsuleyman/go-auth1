FROM golang:1.17.1 AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD go run cmd/main.go --start-service