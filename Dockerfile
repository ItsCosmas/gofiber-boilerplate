# Pull base image
FROM golang:1.24-alpine

# Install git
RUN apk update && apk add --no-cache git

# Install Air for hot reload
RUN go install github.com/air-verse/air@latest