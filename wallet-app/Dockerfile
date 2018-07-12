FROM node:9

RUN mkdir /app
WORKDIR /app
COPY . /app

RUN make install
RUN make build
