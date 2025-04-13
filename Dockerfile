FROM golang:1.24.1-alpine AS builder

RUN apk add --no-cache git

RUN adduser -D -g '' appuser

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cd cmd/orders && CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o run

FROM alpine:3.19

RUN apk add --no-cache curl

COPY --from=builder /etc/passwd /etc/passwd

WORKDIR /app

COPY --from=builder /app/cmd/orders/run .

USER appuser

EXPOSE 8080

ENTRYPOINT ["./run"] 