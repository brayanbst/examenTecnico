FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila el binario principal (main está en cmd/server/main.go)
RUN go build -o qrservice ./cmd/server

FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/qrservice .

# Heroku define PORT, tu código usa os.Getenv("PORT") con default 8080
ENV PORT=8080

CMD ["./qrservice"]
