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
CMD ["yarn", "start", "-p", "80"]

# Port
EXPOSE 80

# Healthcheck
HEALTHCHECK --interval=5s --timeout=2s --retries=12 CMD curl --silent --fail localhost/probe || exit 1