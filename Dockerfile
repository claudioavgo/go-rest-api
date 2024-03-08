FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./goapp .
 
FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/goapp .
EXPOSE 8080
ENTRYPOINT ["./goapp"]