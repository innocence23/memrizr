# Build stage
FROM golang:alpine as builder
WORKDIR /app

COPY . .
RUN go get github.com/cespare/reflex

RUN go build -o memrizr main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/memrizr .
EXPOSE 8080
CMD [ "/app/memrizr" ]
