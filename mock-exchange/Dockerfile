FROM kowalatech/go:1.0.4 as build

WORKDIR /go/src/github.com/kowala-tech/kcoin/
ADD . .

RUN go get -u github.com/golang/dep/cmd/dep && cd mock-exchange && dep ensure --vendor-only
RUN go build -o app mock-exchange/main.go

FROM alpine:3.7
WORKDIR /mock-exchange/
COPY --from=build /go/src/github.com/kowala-tech/kcoin/app .
EXPOSE 8080
ENTRYPOINT ["./app"]
