FROM golang:1.10.3-alpine3.7 as builder
RUN apk update && apk add --update git make gcc musl-dev linux-headers

WORKDIR /faucet/
ADD . .
RUN make faucet

FROM alpine:3.7
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /faucet/
COPY --from=builder /faucet/client/build/bin/faucet .
ADD client/release/run_faucet.sh run_faucet.sh
EXPOSE 80
ENTRYPOINT ["./run_faucet.sh"]
