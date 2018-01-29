FROM golang:1.9.2-alpine as builder
RUN apk update && apk add --update git make gcc musl-dev linux-headers

WORKDIR /bootnode/
ADD . .
RUN make bootnode

FROM alpine:3.7
WORKDIR /bootnode/
COPY --from=builder /bootnode/build/bin/bootnode .
EXPOSE 33445/udp
ENTRYPOINT ["./bootnode"] 
