FROM alpine

LABEL maintainer="me@test.com"

RUN apk add --update nodejs-npm

COPY . /src

RUN npm install

EXPOSE 8080

ENTRYPOINT ["node", "./app.js"]
