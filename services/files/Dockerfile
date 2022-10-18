FROM node:lts-alpine3.15
WORKDIR /files
COPY package*.json ./
COPY . .
EXPOSE 3001
RUN npm install -g npm
RUN npm install
CMD ["npm","run","start"]