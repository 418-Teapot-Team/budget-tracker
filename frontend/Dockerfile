FROM node:lts-alpine

RUN npm install -g serve

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN echo "APP_BASE_API=http://16.16.208.186:8080/" > .env

RUN npm run build

EXPOSE 80
CMD [ "serve", "-s", "-p", "80", "dist" ]
