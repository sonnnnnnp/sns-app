FROM node:22

WORKDIR /frontend

COPY package*.json .

RUN npm ci

COPY . .
