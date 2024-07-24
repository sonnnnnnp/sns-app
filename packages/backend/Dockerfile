FROM node:22

WORKDIR /backend

COPY package*.json .

RUN npm ci

COPY . .
