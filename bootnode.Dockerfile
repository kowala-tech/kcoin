FROM golang:1.10-alpine as builder
RUN apk update && apk add --update git make gcc musl-dev linux-headers

WORKDIR /bootnode/
ADD . .
RUN make bootnode

FROM alpine:latest
WORKDIR /bootnode/
COPY --from=builder /bootnode/build/bin/bootnode .
ENTRYPOINT ["./bootnode"] 
