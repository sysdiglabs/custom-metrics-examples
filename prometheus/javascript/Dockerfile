FROM node:alpine
WORKDIR /app
ADD package.json /app
RUN cd /app && npm install
ADD app.js /app
CMD ["node", "/app/app"]