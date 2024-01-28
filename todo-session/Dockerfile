FROM node:21-alpine3.18
WORKDIR /app
COPY package*.json .
RUN npm install --production
COPY index.js .
EXPOSE 3003
CMD [ "node", "index.js" ]
