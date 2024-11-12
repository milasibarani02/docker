# Use Golang image to build the application
FROM golang:1.20-alpine

WORKDIR /myapp

# Install dependencies and get Go version
RUN go version

# Copy AUTHORS.md before building the app
COPY AUTHORS.md /myapp/AUTHORS.md

# Copy LINKS.md after building the app
COPY LINKS.md /myapp/LINKS.md

# Copy Go source code to the container
COPY . .

# Build the Golang app
RUN go build -o my-go-app .

# Run the app
CMD ["./my-go-app"]
