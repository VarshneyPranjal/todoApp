# Base image (OS)
FROM golang:1.26

# Working directory
WORKDIR /app

# Copy src code to container
COPY . . 

# Run the build commands
RUN go mod download
RUN go build -o todoapp .

# expose port 80
EXPOSE 8080

# serve the app / run the app (keep it running)
CMD ["./todoapp"]
