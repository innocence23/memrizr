# Build stage
FROM golang:alpine as builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
COPY . .
RUN go get github.com/cespare/reflex

RUN go build -o /go/bin/memrizr main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/bin/memrizr .
EXPOSE 8080
CMD [ "/app/memrizr" ]
