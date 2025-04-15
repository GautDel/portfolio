# Use the official Go image as the base image
FROM golang:1.24.1-alpine


# Set the working directory in the container
WORKDIR /app

# Copy the Go binary into the container
COPY . /app

EXPOSE 8090

# Run the Go binary
CMD ["./portfolio", "serve", "--http", "0.0.0.0:8090"]
