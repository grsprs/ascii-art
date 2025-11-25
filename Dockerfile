# Multi-stage build for ascii-art
# Produces a minimal container with just the binary

# Build stage
FROM golang:1.20-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o ascii-art \
    ./cmd/ascii-art

# Final stage
FROM scratch

# Import from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy the binary
COPY --from=builder /build/ascii-art /ascii-art

# Copy banner files
COPY --from=builder /build/banners /banners

# Create non-root user
USER nobody:nobody

# Set entrypoint
ENTRYPOINT ["/ascii-art"]

# Default command
CMD ["--help"]

# Metadata
LABEL maintainer="maintainer@example.com"
LABEL description="ASCII art text renderer"
LABEL version="1.0.0"
LABEL org.opencontainers.image.source="https://github.com/username/ascii-art"
LABEL org.opencontainers.image.documentation="https://github.com/username/ascii-art/blob/main/README.md"
LABEL org.opencontainers.image.licenses="MIT"