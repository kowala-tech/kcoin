FROM golang:1.10-alpine AS build
RUN apk update && apk add --update git make gcc musl-dev linux-headers

WORKDIR /go/src/github.com/kowala-tech/kcoin/wallet-backend/
ADD . .

RUN go build -o backend cmd/main.go

FROM alpine
WORKDIR /backend/
COPY --from=build /go/src/github.com/kowala-tech/kcoin/wallet-backend/backend .
EXPOSE 8080
ENTRYPOINT ["./backend"]
