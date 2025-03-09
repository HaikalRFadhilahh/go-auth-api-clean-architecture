####### COMPILER #######

# Compile Using Base Image Golang
FROM golang:tip-alpine3.21 AS go-app-compiler

# Setting Working Directory
WORKDIR /app

# Copy Source Code
COPY . .

# Install Go Dependencies
RUN go mod tidy
RUN go mod download

# Install Make Command
RUN apk add --no-cache make

# Compile Go Document
RUN make clean
RUN make build

####### RUNTIME #######

# Running Binary Go With Alpine Linux
FROM alpine:3.21.3

# Label
LABEL appname="Go Clean Architecture"
LABEL language="go"
LABEL type="API"
LABEL author="Haikal R Fadhilah"
LABEL os="alpine"

# Setting Working Directory
WORKDIR /app

# Copy Binary File From Builder
COPY --from=go-app-compiler /app/build/main /app

# User Management
RUN addgroup -S goapp && adduser -S -s /bin/sh -h /home/gouser -G goapp gouser

# Chance Owner
RUN chown gouser:goapp -R /app

# ENV
# Static ENV
RUN echo "export APP_HOST=0.0.0.0" >> /home/gouser/.profile 
RUN echo "export APP_PORT=8000" >> /home/gouser/.profile

# Dynamic ENV
ENV DB_HOST=127.0.0.1
ENV DB_PORT=3306
ENV DB_USERNAME=root
ENV DB_PASSWORD=
ENV DB_NAME=

ENV JWT_SECRET=
ENV JWT_EXPIRED_HOUR=1

# Health Check
HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD curl -X GET http://127.0.0.1:8000/health || exit 1

# Export Port
EXPOSE 8000

# Change User Runtime
USER gouser

# Running Go
CMD [ "sh","/app/main" ]