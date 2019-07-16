FROM node:12.6

RUN mkdir -p /home/website
WORKDIR /home/website

# Get Deps
COPY yarn.lock .
COPY package.json .
RUN yarn

# Build server
COPY components/ /home/website/components
COPY pages/ /home/website/pages
RUN yarn build

# Start Server
EXPOSE 80
CMD ["yarn", "start", "-p", "80"]
