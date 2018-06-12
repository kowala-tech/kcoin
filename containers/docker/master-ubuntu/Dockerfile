FROM ubuntu:xenial

ENV PATH=/usr/lib/go-1.9/bin:$PATH

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang-1.9 git make gcc libc-dev ca-certificates && \
  git clone --depth 1 --branch release/1 https://github.com/kowala-tech/kcoin && \
  (cd kcoin && make kcoin) && \
  cp kcoin/build/bin/kcoin /kcoin && \
  apt-get remove -y golang git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /kcoin

EXPOSE 11223
EXPOSE 22334

ENTRYPOINT ["/kcoin"]
