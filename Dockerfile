# Pull base image
FROM golang:1.16-alpine

# Install git
RUN apk update && apk add --no-cache git

# Install Air for hot reload
RUN go get -u github.com/cosmtrek/air