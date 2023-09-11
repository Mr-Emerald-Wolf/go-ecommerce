# Build stage
FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o main cmd/server/main.go

# Run stage
FROM alpine

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .

CMD [ "/app/main" ]