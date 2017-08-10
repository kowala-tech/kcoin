FROM ubuntu:xenial

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang git make gcc libc-dev ca-certificates && \
  git clone --depth 1 https://github.com/kowala-tech/kUSD && \
  (cd kUSD && make geth) && \
  cp kUSD/build/bin/geth /geth && \
  apt-get remove -y golang git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /kUSD

EXPOSE 11223
EXPOSE 22334

ENTRYPOINT ["/geth"]
