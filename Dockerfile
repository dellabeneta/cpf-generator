FROM golang:alpine3.22 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o cpf-generator

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/cpf-generator .

COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./cpf-generator"]