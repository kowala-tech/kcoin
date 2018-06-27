FROM golang:1.9.2-alpine as builder
WORKDIR /go/src/github.com/kowala-tech/kcoin
RUN apk update; apk add --no-cache git curl alpine-sdk
COPY . .
RUN go build -a -o app notifications/cmd/transactions_publisher/main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/kowala-tech/kcoin/app .
CMD ["./app"] 
