# syntax=base image i used
FROM node:1.21-alpine

WORKDIR /usr/src/app

COPY /web/* .

RUN npm install

RUN npm build

EXPOSE 3000

CMD ["npm", "start"]