FROM golang:1.10.3-alpine3.7 as builder
RUN apk update && apk add --update git make gcc musl-dev linux-headers

WORKDIR /bootnode/
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
COPY --from=builder /bootnode/client/build/bin/bootnode .
ENTRYPOINT ["./bootnode"] 
