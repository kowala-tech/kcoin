FROM golang:1.10.3-alpine3.7 as builder
RUN apk update && apk add --update git make gcc musl-dev linux-headers

WORKDIR /bootnode/
ADD . .
RUN make bootnode

FROM alpine:3.7
WORKDIR /bootnode/
COPY --from=builder /bootnode/client/build/bin/bootnode .
ENTRYPOINT ["./bootnode"] 
