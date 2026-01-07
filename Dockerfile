# build stage
FROM golang:1.25 AS builder

WORKDIR /app

# install C toolchain for CGO (required by go-sqlite3)
RUN apt-get update && apt-get install -y \
    gcc \
    libc6-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# build the correct path
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api

# runtime stage
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/app .

# creates directory for SQLite DB
WORKDIR /data

EXPOSE 8080

CMD ["/app/app"]