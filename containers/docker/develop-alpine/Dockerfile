FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 https://github.com/kowala-tech/kUSD && \
  (cd kUSD && make kusd) && \
  cp kUSD/build/bin/kusd /kusd && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /kUSD && rm -rf /var/cache/apk/*

EXPOSE 11223
EXPOSE 22334

ENTRYPOINT ["/kusd"]
