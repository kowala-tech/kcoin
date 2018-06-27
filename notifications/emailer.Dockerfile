FROM golang:1.9.2-alpine as builder
WORKDIR /go/src/github.com/kowala-tech/kcoin/notifications/
RUN apk update; apk add --no-cache git curl alpine-sdk
COPY . .
RUN go build -a -o app cmd/emailer/main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/kowala-tech/kcoin/notifications/app .
CMD ["./app"] 
