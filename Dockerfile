# Start from golang base image
FROM golang:alpine as build

# Add Maintainer info
LABEL maintainer="Markus Krebs"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /build

FROM alpine:latest

# Set working directory
WORKDIR /

# Copy compiled binary
COPY --from=build /build /build
COPY --from=build /app/seeds /seeds

# Expose port 8080 to the outside world
EXPOSE 8080

# Set user as non root
USER nonroot:nonroot

# Run the executable
CMD [ "/build" ]