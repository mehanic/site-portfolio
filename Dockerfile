# Use an official Golang image to build the app
FROM docker.io/golang:1.23 AS builder

WORKDIR /app

# Copy go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o site-portfolio cmd/main.go

# Use a minimal base image for running the app
FROM debian:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/site-portfolio /app/site-portfolio

# Copy static files & templates
COPY static/ static/
COPY templates/ templates/

# Copy the environment file
COPY .env .env

# Set environment variables (adjust as needed)
ENV PORT=8080
ENV DB_HOST=db
ENV DB_USER=postgres
ENV DB_PASSWORD=your-secure-password
ENV DB_NAME=site_db
ENV DB_PORT=5432

# Expose the app port
EXPOSE 8080

# Run the application
CMD ["/app/site-portfolio"]
