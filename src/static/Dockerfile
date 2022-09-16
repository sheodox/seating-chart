FROM node:16 AS dev
WORKDIR /usr/src/app

ENV NODE_ENV=development
COPY package*.json ./
RUN npm ci
CMD npm run dev -- --host 0.0.0.0

