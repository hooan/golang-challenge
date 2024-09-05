# Use the official Golang image as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the source code from the current directory to the working directory inside the container
COPY . .

# Build the Go application
RUN make build

# Expose the port on which the application will run
EXPOSE 8080

# Set the command to run the executable when the container starts
CMD ["make","run"]