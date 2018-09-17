FROM node:9.5-alpine
RUN apk add --update git python make g++ libnotify openssl
RUN npm install -g gulp-cli

# SSL self-signed certificate for localhost.
RUN openssl genrsa -des3 -passout pass:x -out server.pass.key 2048 && \
    openssl rsa -passin pass:x -in server.pass.key -out server.key && \
    openssl req -new -key server.key -out server.csr -subj "/C=US/ST=California/L=California/O=localhost/OU=localhost/CN=localhost" && \
    openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -out server.crt

WORKDIR /wallet-tools

CMD ["npm","help"]
