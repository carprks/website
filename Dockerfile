######### Backend Build #########
FROM golang:1.12 AS backend_build
RUN mkdir -p /home/main
WORKDIR /home/main

# Dependencies
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download

# ENV
ARG SERVICE_NAME=website
ARG SERVICE_DEPENDENCIES

# Copy the rest
COPY main.go .
COPY backend backend

# Build
ARG build
ARG version
RUN CGO_ENABLED=0 go build -ldflags "-w -s -X github.com/carprks/website/backend/website.Version=${version} -X github.com/carprks/website/backend/website.Build=${build}" -o ${SERVICE_NAME} .
RUN cp ${SERVICE_NAME} /

######### Frontend Build #######
FROM node:12.8.0 AS frontend_build
RUN mkdir -p /home/frontend
WORKDIR /home/frontend
COPY frontend /home/frontend
RUN yarn install
RUN yarn tailwind build css/prebuild.css -o css/tailwind.css

######### Distribution #########
FROM alpine
ARG SERVICE_NAME=website
RUN apk update && apk upgrade
RUN apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
RUN apk add curl
RUN rm -rf /var/cache/apk/*

# Move from builds
COPY --from=backend_build /${SERVICE_NAME} /home/
COPY frontend /home/frontend
COPY --from=frontend_build /home/frontend/css /home/frontend/css

# Set Timezone
ENV TZ=Europe/London

# EntryPoint Create
WORKDIR /home
ENV _SERVICENAME=${SERVICE_NAME}
RUN echo "#!/usr/bin/env bash" > ./entrypoint.sh
RUN echo "./${SERVICE_NAME}" >> ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

# EntryPoint
ENTRYPOINT ["sh", "./entrypoint.sh"]

# HealthCheck
HEALTHCHECK --interval=5s --timeout=2s --retries=12 CMD curl --silent --fail localhost/probe || exit 1

# Port
EXPOSE 80
