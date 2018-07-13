FROM node:latest

COPY explorer/package*.json ./

RUN npm install

COPY explorer/. .

EXPOSE 3000

CMD [ "node", "tools/sync.js" ]
