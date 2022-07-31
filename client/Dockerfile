FROM node:18 as Dev

WORKDIR /usr/app

ENV LANG=C.UTF-8 \
  TZ=Asia/Tokyo

COPY package.json yarn.lock ./
COPY tsconfig*.json ./
COPY vite.config.ts ./
COPY index.html ./

RUN yarn install

EXPOSE 5173