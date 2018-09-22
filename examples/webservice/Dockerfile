# Stage 1: get sources from npm
FROM node:alpine AS sources
WORKDIR /app/
COPY package*.json /app/
RUN npm install

# Stage 2: build minified production code
FROM node:alpine AS intermediate
WORKDIR /app/
COPY --from=sources /app/ /app/
COPY . /app/

# Stage 3
FROM node:alpine AS production
WORKDIR /app/
COPY --from=intermediate /app/ /app/
ENV NPM_CONFIG_LOGLEVEL warn
CMD npm start
EXPOSE 8080