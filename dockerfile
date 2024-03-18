FROM golang:1.22-alpine AS builder

# Define the desired project address inside the container
WORKDIR /go/src/app

# Copy the source code into the container
COPY . .

# Build the project with cross-compilation for Linux/ARM64
RUN GOOS=linux GOARCH=arm64 go build -o /go/bin/app

# Start a new stage to create a smaller final image
FROM alpine:latest

# Copy the built binary from the previous stage
COPY --from=builder /go/bin/app /usr/local/bin/app

# Copy the endpointList.json file into the final image
COPY --from=builder /go/src/app/src/endpointList.json /usr/local/bin/src/endpointList.json

# Expose the port
EXPOSE 3000

# Run the built binary
CMD ["app"]
