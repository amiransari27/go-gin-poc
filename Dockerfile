# Build stage for React
FROM node:22-alpine AS frontend-build
WORKDIR /app/client
COPY client/package.json ./
COPY client/yarn.lock ./
RUN yarn install
COPY client/ ./
RUN yarn run build

# Build stage for Go
FROM golang:1.24.3-bookworm  AS backend-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .


# Final stage
FROM alpine:latest
WORKDIR /app
# Copy the Go binary
COPY --from=backend-build /app/main .
# Copy the React build
COPY --from=frontend-build /app/client/dist ./client/dist

# Expose port
EXPOSE 3002

# Run the binary
CMD ["./main"]