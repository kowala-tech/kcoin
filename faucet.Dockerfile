FROM golang:1.10-alpine as builder
RUN apk update && apk add --update git make gcc musl-dev linux-headers

WORKDIR /faucet/
ADD . .
RUN make faucet

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /faucet/
COPY --from=builder /faucet/build/bin/faucet .
ADD release/testnet_genesis.json genesis.json
ADD release/run_faucet.sh run_faucet.sh
EXPOSE 80
ENTRYPOINT ["./run_faucet.sh"]
