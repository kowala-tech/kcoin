FROM golang:1.9.2-alpine as builder
RUN apk update && apk upgrade && apk add --update git make gcc musl-dev linux-headers util-linux pciutils usbutils coreutils binutils findutils grep bash bash-doc bash-completion bash bash-doc bash-completion util-linux pciutils usbutils coreutils binutils findutils grep build-base gcc abuild binutils binutils-doc gcc-doc man man-pages mdocml-apropos less less-doc

WORKDIR /bootnode/
ADD . .
RUN make bootnode

FROM alpine:3.7
WORKDIR /bootnode/
COPY --from=builder /bootnode/build/bin/bootnode .
ENTRYPOINT ["./bootnode"]
