FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/api/main.go

# Stage final
FROM alpine:latest

WORKDIR /app
COPY --from=build /app/server .

EXPOSE 4000

CMD ["./server"]
