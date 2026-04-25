FROM node:22-alpine

WORKDIR /app

COPY package*.json ./

RUN RUN npm ci --only=production

COPY . .

RUN npm run build

RUN npm prune --production

EXPOSE 3000