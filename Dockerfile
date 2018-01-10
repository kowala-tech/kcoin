FROM alpine:3.7

ADD . /kusd
rUN \
  apk update && apk add --update git go make gcc musl-dev linux-headers && \
  go version && \
  (cd /kusd && make kusd)                           && \
  cp /kusd/build/bin/kusd /usr/local/bin/           && \
  apk del git go make gcc musl-dev linux-headers          && \
  rm -rf /kusd && rm -rf /var/cache/apk/*

EXPOSE 11223
EXPOSE 22334
EXPOSE 22334/udp

ENTRYPOINT ["kusd"]
