# Node 12.11-alpine is our base image for building & testing
FROM node:12.22-alpine

RUN apk add --no-cache ca-certificates git

# npm packages version
ARG ganache_version="6.11.0"
ARG truffle_version="5.4.0"

# install node packages
RUN npm -g config set user root
RUN npm install -g ganache-cli@${ganache_version} truffle@${truffle_version}
