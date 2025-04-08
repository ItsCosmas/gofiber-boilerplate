# Pull base image
FROM golang:1.24-alpine

# Install git and # Install Air for hot reload
RUN apk update && apk add --no-cache git && go install github.com/air-verse/air@latest