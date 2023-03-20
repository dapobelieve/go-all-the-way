FROM golang:1.19-alpine3.16 AS builder

WORKDIR /app

COPY . .

COPY go.mod go.sum ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Path: Dockerfile
FROM alpine:3.16

WORKDIR /root/

COPY --from=builder app/main .

EXPOSE 8080

CMD ["./main"]