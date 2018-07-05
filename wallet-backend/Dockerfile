FROM kowalatech/go:1.0.4 as build

WORKDIR /go/src/github.com/kowala-tech/kcoin/
ADD . .

RUN go get -u github.com/golang/dep/cmd/dep && cd wallet-backend && dep ensure --vendor-only
RUN go build -o app wallet-backend/cmd/main.go

FROM alpine:3.7
WORKDIR /backend/
COPY --from=build /go/src/github.com/kowala-tech/kcoin/app .
EXPOSE 8080
ENTRYPOINT ["./app"]
