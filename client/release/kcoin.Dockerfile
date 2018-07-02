FROM golang:1.10.3-alpine3.7 as builder
RUN apk update && apk add --update git make gcc musl-dev linux-headers

WORKDIR /kcoin/
ADD . .

ARG CI
ARG DRONE
ARG DRONE_REPO
ARG DRONE_COMMIT_SHA
ARG DRONE_COMMIT_BRANCH
ARG DRONE_TAG
ARG DRONE_BUILD_NUMBER
ARG DRONE_BUILD_EVENT

RUN make kcoin control

FROM alpine:3.7
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /kcoin/
COPY --from=builder /kcoin/client/build/bin/kcoin .
COPY --from=builder /kcoin/client/build/bin/control .
EXPOSE 22334
EXPOSE 22334/udp
EXPOSE 8080
ADD client/release/kcoin.sh .
ENTRYPOINT ["./kcoin.sh"]
RUN mkdir -p /root/.kcoin/kusd/keystore
