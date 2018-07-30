FROM node
WORKDIR /app
COPY package.json /app
COPY app.js /app
RUN cd /app && npm install
CMD ["node", "/app/app"]