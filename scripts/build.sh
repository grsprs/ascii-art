#!/bin/bash

# Build script for ascii-art project
# Usage: ./scripts/build.sh [options]

set -euo pipefail

# Configuration
PROJECT_NAME="ascii-art"
CMD_PATH="./cmd/ascii-art"
DIST_DIR="dist"
GO_VERSION="1.20"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Functions
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

check_go_version() {
    if ! command -v go &> /dev/null; then
        log_error "Go is not installed"
        exit 1
    fi
    
    local go_version=$(go version | cut -d' ' -f3 | sed 's/go//')
    local required_version="1.20"
    
    if ! printf '%s\n%s\n' "$required_version" "$go_version" | sort -V -C; then
        log_warn "Go version $go_version detected, $required_version+ recommended"
    fi
}

clean() {
    log_info "Cleaning build artifacts..."
    rm -rf "$DIST_DIR"
    go clean -cache
    go clean -modcache
}

test() {
    log_info "Running tests..."
    go test -race -coverprofile=coverage.out ./...
    
    log_info "Running golden tests..."
    go test -run Golden ./...
    
    log_info "Generating coverage report..."
    go tool cover -html=coverage.out -o coverage.html
}

lint() {
    log_info "Running linters..."
    
    # Format check
    if [ "$(gofmt -s -l . | wc -l)" -gt 0 ]; then
        log_error "Code is not formatted. Run: gofmt -s -w ."
        gofmt -s -l .
        exit 1
    fi
    
    # Go vet
    go vet ./...
    
    # staticcheck (if available)
    if command -v staticcheck &> /dev/null; then
        staticcheck ./...
    else
        log_warn "staticcheck not found, skipping static analysis"
    fi
}

build_single() {
    local goos=$1
    local goarch=$2
    local version=${3:-"dev"}
    
    local ext=""
    if [ "$goos" = "windows" ]; then
        ext=".exe"
    fi
    
    local binary_name="${PROJECT_NAME}-${goos}-${goarch}${ext}"
    local output_path="${DIST_DIR}/${binary_name}"
    
    log_info "Building $binary_name..."
    
    GOOS=$goos GOARCH=$goarch go build \
        -ldflags="-s -w -X main.version=$version" \
        -o "$output_path" \
        "$CMD_PATH"
    
    # Generate checksum
    if command -v sha256sum &> /dev/null; then
        (cd "$DIST_DIR" && sha256sum "$binary_name" > "${binary_name}.sha256")
    elif command -v shasum &> /dev/null; then
        (cd "$DIST_DIR" && shasum -a 256 "$binary_name" > "${binary_name}.sha256")
    fi
    
    log_info "Built: $output_path"
}

build_all() {
    local version=${1:-"dev"}
    
    log_info "Building all platforms..."
    mkdir -p "$DIST_DIR"
    
    # Build matrix
    local platforms=(
        "linux amd64"
        "linux arm64"
        "darwin amd64"
        "darwin arm64"
        "windows amd64"
    )
    
    for platform in "${platforms[@]}"; do
        local goos=$(echo $platform | cut -d' ' -f1)
        local goarch=$(echo $platform | cut -d' ' -f2)
        build_single "$goos" "$goarch" "$version"
    done
    
    log_info "Builds completed in $DIST_DIR/"
}

package() {
    log_info "Creating release packages..."
    
    cd "$DIST_DIR"
    
    for file in ascii-art-*; do
        if [[ ! "$file" =~ \.(sha256|zip|tar\.gz)$ ]]; then
            if [[ "$file" =~ windows ]]; then
                zip "${file%.exe}.zip" "$file" "${file}.sha256"
                log_info "Created ${file%.exe}.zip"
            else
                tar -czf "${file}.tar.gz" "$file" "${file}.sha256"
                log_info "Created ${file}.tar.gz"
            fi
        fi
    done
    
    cd ..
}

install_local() {
    log_info "Installing locally..."
    go install "$CMD_PATH"
    log_info "Installed to $(go env GOPATH)/bin/$PROJECT_NAME"
}

benchmark() {
    log_info "Running benchmarks..."
    go test -bench=. -benchmem ./...
}

security_scan() {
    log_info "Running security scans..."
    
    # Vulnerability check
    if command -v govulncheck &> /dev/null; then
        govulncheck ./...
    else
        log_warn "govulncheck not found, install with: go install golang.org/x/vuln/cmd/govulncheck@latest"
    fi
    
    # Security scan
    if command -v gosec &> /dev/null; then
        gosec ./...
    else
        log_warn "gosec not found, install with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"
    fi
}

usage() {
    cat << EOF
Usage: $0 [command] [options]

Commands:
    build [version]     Build for current platform (default: dev)
    build-all [version] Build for all platforms (default: dev)
    test               Run all tests
    lint               Run linters and format checks
    clean              Clean build artifacts
    package            Create release packages
    install            Install locally
    benchmark          Run performance benchmarks
    security           Run security scans
    help               Show this help

Examples:
    $0 build v1.0.0
    $0 build-all v1.2.3
    $0 test
    $0 lint

EOF
}

# Main execution
main() {
    check_go_version
    
    case "${1:-help}" in
        "build")
            build_single "$(go env GOOS)" "$(go env GOARCH)" "${2:-dev}"
            ;;
        "build-all")
            build_all "${2:-dev}"
            ;;
        "test")
            test
            ;;
        "lint")
            lint
            ;;
        "clean")
            clean
            ;;
        "package")
            package
            ;;
        "install")
            install_local
            ;;
        "benchmark")
            benchmark
            ;;
        "security")
            security_scan
            ;;
        "help"|"--help"|"-h")
            usage
            ;;
        *)
            log_error "Unknown command: $1"
            usage
            exit 1
            ;;
    esac
}

main "$@"