FROM golang:1.9.2-alpine as builder
RUN apk update && apk add --update git make gcc musl-dev linux-headers

WORKDIR /faucet/
ADD . .
RUN make faucet

FROM alpine:3.7
WORKDIR /faucet/
COPY --from=builder /faucet/build/bin/faucet .
EXPOSE 80
ENTRYPOINT ["./faucet"]
