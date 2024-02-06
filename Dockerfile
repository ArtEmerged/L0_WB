FROM golang:1.21.6-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o subscriber ./cmd/app/main.go

FROM golang:1.21.6-alpine

WORKDIR /app

COPY --from=builder /app /app

CMD ["./subscriber"]
