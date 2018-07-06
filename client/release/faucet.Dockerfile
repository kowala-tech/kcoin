FROM kowalatech/go:1.0.4 as builder

WORKDIR /go/src/github.com/kowala-tech/kcoin/
ADD . .

ARG CI
ARG DRONE
ARG DRONE_REPO
ARG DRONE_COMMIT_SHA
ARG DRONE_COMMIT_BRANCH
ARG DRONE_TAG
ARG DRONE_BUILD_NUMBER
ARG DRONE_BUILD_EVENT

RUN make faucet

FROM alpine:3.7
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /faucet/
COPY --from=builder /go/src/github.com/kowala-tech/kcoin/client/build/bin/faucet .
ADD client/release/run_faucet.sh run_faucet.sh
EXPOSE 80
ENTRYPOINT ["./run_faucet.sh"]
