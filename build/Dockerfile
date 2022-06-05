# build stage
FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go build -v -o app-go .

# distribute stage
FROM alpine:latest
WORKDIR /root
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app-go .
ENTRYPOINT ./app-go
