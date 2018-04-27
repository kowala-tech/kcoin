FROM golang:1.9.2-alpine as builder
RUN apk update && apk upgrade && apk add --update git make gcc musl-dev linux-headers util-linux pciutils usbutils coreutils binutils findutils grep bash bash-doc bash-completion bash bash-doc bash-completion util-linux pciutils usbutils coreutils binutils findutils grep build-base gcc abuild binutils binutils-doc gcc-doc man man-pages mdocml-apropos less less-doc

WORKDIR /kcoin/
ADD . .
RUN make kcoin

FROM alpine:3.7
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /kcoin/
COPY --from=builder /kcoin/build/bin/kcoin .
EXPOSE 22334
EXPOSE 22334/udp
ADD release/kcoin_with_new_account.sh .
ADD release/testnet_console.toml .
ADD release/testnet_genesis.json genesis.json
ENTRYPOINT ["./kcoin_with_new_account.sh"]
CMD ["--config", "/kcoin/testnet_console.toml"]
