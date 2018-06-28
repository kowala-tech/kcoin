FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 https://github.com/kowala-tech/kcoin/client && \
  (cd kcoin && make kcoin) && \
  cp kcoin/build/bin/kcoin /kcoin && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /kcoin && rm -rf /var/cache/apk/*

EXPOSE 11223
EXPOSE 22334

ENTRYPOINT ["/kcoin"]
