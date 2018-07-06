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

RUN make bootnode

FROM alpine:3.7
WORKDIR /bootnode/
COPY --from=builder /go/src/github.com/kowala-tech/kcoin/client/build/bin/bootnode .
ENTRYPOINT ["./bootnode"] 
