FROM alpine:3.5

ADD . /kUSD
RUN \
  apk add --update git go make gcc musl-dev linux-headers && \
  (cd kUSD && make geth)                           && \
  cp kUSD/build/bin/geth /usr/local/bin/           && \
  apk del git go make gcc musl-dev linux-headers          && \
  rm -rf /kUSD && rm -rf /var/cache/apk/*

EXPOSE 11223
EXPOSE 22334
EXPOSE 22334/udp

ENTRYPOINT ["geth"]
