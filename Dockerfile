# Multi-stage build for React + Go application

# Stage 1: Build React frontend
FROM node:25-alpine AS frontend-build
WORKDIR /app/frontend
COPY front-end/package*.json ./
RUN npm ci
COPY front-end/ ./
RUN npm run build

# Stage 2: Build Go backend
FROM golang:1.25-alpine AS backend-build
WORKDIR /app/backend

# Copy go mod files first for better caching
COPY back-end/go.mod back-end/go.sum ./
RUN go mod download

# Copy source code
COPY back-end/ ./

# Build with cache mount for faster rebuilds
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 3: Final runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

# Copy the Go binary
COPY --from=backend-build /app/backend/main .

# Copy the React build files
COPY --from=frontend-build /app/frontend/dist ./static

# Expose port
EXPOSE 8080

# Run the Go server
CMD ["./main"]