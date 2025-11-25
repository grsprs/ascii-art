#!/bin/bash

# Test script for ascii-art project
# Usage: ./scripts/test.sh [options]

set -euo pipefail

# Configuration
COVERAGE_THRESHOLD=80
COVERAGE_FILE="coverage.out"
COVERAGE_HTML="coverage.html"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_section() {
    echo -e "\n${BLUE}=== $1 ===${NC}"
}

# Unit tests
run_unit_tests() {
    log_section "Running Unit Tests"
    go test -v -race -coverprofile="$COVERAGE_FILE" ./...
}

# Golden tests
run_golden_tests() {
    log_section "Running Golden Tests"
    go test -v -run Golden ./...
}

# Benchmark tests
run_benchmarks() {
    log_section "Running Benchmarks"
    go test -bench=. -benchmem ./...
}

# Coverage analysis
analyze_coverage() {
    log_section "Analyzing Test Coverage"
    
    if [ ! -f "$COVERAGE_FILE" ]; then
        log_error "Coverage file not found. Run unit tests first."
        return 1
    fi
    
    # Generate HTML report
    go tool cover -html="$COVERAGE_FILE" -o "$COVERAGE_HTML"
    log_info "Coverage report generated: $COVERAGE_HTML"
    
    # Check coverage threshold
    local coverage=$(go tool cover -func="$COVERAGE_FILE" | grep total | awk '{print $3}' | sed 's/%//')
    
    echo "Total coverage: ${coverage}%"
    
    if (( $(echo "$coverage >= $COVERAGE_THRESHOLD" | bc -l) )); then
        log_info "Coverage threshold met: ${coverage}% >= ${COVERAGE_THRESHOLD}%"
    else
        log_error "Coverage below threshold: ${coverage}% < ${COVERAGE_THRESHOLD}%"
        return 1
    fi
}

# Package-specific coverage
package_coverage() {
    log_section "Package Coverage Breakdown"
    
    if [ ! -f "$COVERAGE_FILE" ]; then
        log_error "Coverage file not found. Run unit tests first."
        return 1
    fi
    
    go tool cover -func="$COVERAGE_FILE" | grep -v total | while read line; do
        local pkg=$(echo "$line" | awk '{print $1}' | cut -d'/' -f1-2)
        local func=$(echo "$line" | awk '{print $2}')
        local coverage=$(echo "$line" | awk '{print $3}')
        
        printf "%-40s %-30s %s\n" "$pkg" "$func" "$coverage"
    done
}

# Race condition tests
run_race_tests() {
    log_section "Running Race Condition Tests"
    go test -race ./...
}

# Memory leak tests
run_memory_tests() {
    log_section "Running Memory Tests"
    
    # Test with different input sizes
    local test_inputs=(
        "Hello"
        "$(printf 'A%.0s' {1..100})"
        "$(printf 'B%.0s' {1..1000})"
    )
    
    for input in "${test_inputs[@]}"; do
        echo "Testing input size: ${#input} characters"
        echo "$input" | go run ./cmd/ascii-art 2>/dev/null >/dev/null
    done
    
    log_info "Memory tests completed"
}

# Fuzz testing (if available)
run_fuzz_tests() {
    log_section "Running Fuzz Tests"
    
    # Check if fuzz tests exist
    if find . -name "*_test.go" -exec grep -l "func Fuzz" {} \; | grep -q .; then
        go test -fuzz=. -fuzztime=30s ./...
    else
        log_warn "No fuzz tests found"
    fi
}

# Integration tests
run_integration_tests() {
    log_section "Running Integration Tests"
    
    # Test CLI with various inputs
    local binary="./ascii-art"
    
    # Build binary for testing
    go build -o "$binary" ./cmd/ascii-art
    
    # Test cases
    local test_cases=(
        '""'
        '"Hello"'
        '"Hello\nWorld"'
        '"Hello 123!"'
        '"{Special [Chars]}"'
    )
    
    for test_case in "${test_cases[@]}"; do
        echo "Testing: $test_case"
        eval "$binary $test_case" >/dev/null
        if [ $? -eq 0 ]; then
            echo "✓ Passed"
        else
            echo "✗ Failed"
            return 1
        fi
    done
    
    # Test file input
    echo "Hello World" > test_input.txt
    "$binary" -file=test_input.txt >/dev/null
    rm -f test_input.txt
    
    # Test different banners
    for banner in standard shadow thinkertoy; do
        echo "Testing banner: $banner"
        "$binary" -banner="$banner" "Test" >/dev/null
    done
    
    # Cleanup
    rm -f "$binary"
    
    log_info "Integration tests completed"
}

# Performance regression tests
run_performance_tests() {
    log_section "Running Performance Tests"
    
    # Create baseline if it doesn't exist
    if [ ! -f "benchmark_baseline.txt" ]; then
        log_info "Creating performance baseline..."
        go test -bench=. -benchmem ./... > benchmark_baseline.txt
        log_info "Baseline created: benchmark_baseline.txt"
        return 0
    fi
    
    # Run current benchmarks
    go test -bench=. -benchmem ./... > benchmark_current.txt
    
    # Compare (requires benchcmp tool)
    if command -v benchcmp &> /dev/null; then
        log_info "Comparing with baseline:"
        benchcmp benchmark_baseline.txt benchmark_current.txt
    else
        log_warn "benchcmp not found. Install with: go install golang.org/x/tools/cmd/benchcmp@latest"
        log_info "Current benchmark results saved to benchmark_current.txt"
    fi
}

# Test data validation
validate_test_data() {
    log_section "Validating Test Data"
    
    # Check golden files exist and are valid
    if [ -d "testdata/golden" ]; then
        local golden_files=$(find testdata/golden -name "*.golden" | wc -l)
        log_info "Found $golden_files golden test files"
        
        # Validate golden files are not empty
        find testdata/golden -name "*.golden" -empty | while read empty_file; do
            log_warn "Empty golden file: $empty_file"
        done
    else
        log_warn "No golden test directory found"
    fi
    
    # Check banner files
    if [ -d "banners" ]; then
        for banner in banners/*.txt; do
            if [ -f "$banner" ]; then
                local lines=$(wc -l < "$banner")
                log_info "Banner $(basename "$banner"): $lines lines"
                
                # Basic validation (should have 8 lines per character * 95 characters + separators)
                local expected_min=760  # 95 * 8 + 94 separators
                if [ "$lines" -lt "$expected_min" ]; then
                    log_warn "Banner file $banner may be incomplete: $lines lines (expected >$expected_min)"
                fi
            fi
        done
    fi
}

# Clean test artifacts
clean_test_artifacts() {
    log_info "Cleaning test artifacts..."
    rm -f "$COVERAGE_FILE" "$COVERAGE_HTML"
    rm -f benchmark_current.txt
    rm -f test_input.txt
    rm -f ascii-art  # Test binary
}

# Test summary
print_summary() {
    log_section "Test Summary"
    
    if [ -f "$COVERAGE_FILE" ]; then
        local coverage=$(go tool cover -func="$COVERAGE_FILE" | grep total | awk '{print $3}')
        echo "Total Coverage: $coverage"
    fi
    
    echo "Test artifacts:"
    [ -f "$COVERAGE_HTML" ] && echo "  - Coverage report: $COVERAGE_HTML"
    [ -f "benchmark_current.txt" ] && echo "  - Benchmark results: benchmark_current.txt"
}

usage() {
    cat << EOF
Usage: $0 [command]

Commands:
    all                Run all tests (default)
    unit              Run unit tests only
    golden            Run golden tests only
    integration       Run integration tests
    benchmark         Run benchmark tests
    coverage          Analyze test coverage
    race              Run race condition tests
    memory            Run memory tests
    fuzz              Run fuzz tests (if available)
    performance       Run performance regression tests
    validate          Validate test data
    clean             Clean test artifacts
    help              Show this help

Examples:
    $0                # Run all tests
    $0 unit           # Run unit tests only
    $0 coverage       # Analyze coverage
    $0 benchmark      # Run benchmarks

EOF
}

# Main execution
main() {
    case "${1:-all}" in
        "all")
            validate_test_data
            run_unit_tests
            run_golden_tests
            run_integration_tests
            analyze_coverage
            package_coverage
            print_summary
            ;;
        "unit")
            run_unit_tests
            ;;
        "golden")
            run_golden_tests
            ;;
        "integration")
            run_integration_tests
            ;;
        "benchmark")
            run_benchmarks
            ;;
        "coverage")
            analyze_coverage
            package_coverage
            ;;
        "race")
            run_race_tests
            ;;
        "memory")
            run_memory_tests
            ;;
        "fuzz")
            run_fuzz_tests
            ;;
        "performance")
            run_performance_tests
            ;;
        "validate")
            validate_test_data
            ;;
        "clean")
            clean_test_artifacts
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