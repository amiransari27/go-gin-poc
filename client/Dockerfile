FROM node:22-alpine

WORKDIR /app

# Add environment variables to ensure hot reloading works
ENV CHOKIDAR_USEPOLLING=true
ENV WATCHPACK_POLLING=true
ENV WDS_SOCKET_PORT=0

COPY package.json ./
COPY yarn.lock ./

RUN yarn install

COPY . ./

EXPOSE 3000

CMD ["npm", "run", "dev"]