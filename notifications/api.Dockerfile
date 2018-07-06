FROM kowalatech/go:1.0.4 as builder
WORKDIR /go/src/github.com/kowala-tech/kcoin
COPY . .
RUN cd notifications && dep ensure --vendor-only
RUN go build -a -o app notifications/cmd/api/main.go

FROM alpine:latest3.7
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/kowala-tech/kcoin/app .
CMD ["./app"] 
