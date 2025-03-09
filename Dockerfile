# Stage 1: Build stage
FROM golang:1.24.0-alpine AS builder

# Install git and required dependencies
RUN apk update && apk add --no-cache git

# Set the working directory
WORKDIR /src

# Set GOPROXY environment variable
ENV GOPROXY=https://proxy.golang.org,direct

# Define build arguments (with defaults)
ARG TARGETOS
ARG TARGETARCH

# Set default values for when using `docker build`
RUN export TARGETOS=${TARGETOS:-$(go env GOOS)} && \
    export TARGETARCH=${TARGETARCH:-$(go env GOARCH)}

# Copy the application source code
COPY . .

# Download dependencies
RUN go mod download

# Build the Go application for the target OS and architecture
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -installsuffix cgo -o main cmd/app/main.go

# Stage 2: Final stage
FROM alpine:3.21.3 AS production

# Install tzdata to handle timezones
RUN apk update && apk add --no-cache tzdata

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /src/main .
COPY --from=builder /src/.env .
COPY --from=builder /src/logs ./logs

# Optionally set timezone (if you want the container to run in Asia/Jakarta time)
ENV TZ=Asia/Jakarta

# Set the default port value using ARG and ENV
ARG HTTP_PORT=8080
ENV HTTP_PORT=${HTTP_PORT}

# Dynamically expose port (optional, but EXPOSE is not mandatory for the app to work)
EXPOSE ${HTTP_PORT}

# Command to run the executable
CMD ["./main"]
