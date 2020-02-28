FROM golang:alpine AS builder
WORKDIR /app
COPY . /app
RUN go build -o main .

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT ["./main"]