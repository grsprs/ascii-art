# Contributing

Thanks for your interest in contributing to this project. Here's what you need to know to get started.

## Developer Setup

### Prerequisites

- Go 1.20 or later
- Git
- Make (optional, for convenience scripts)

### Local Development

```bash
# Clone repository
git clone https://github.com/spirosnikoloudakis/ascii-art.git
cd ascii-art

# Verify setup
go version
go env GOPATH

# Run tests
go test ./...

# Run locally
go run ./cmd/ascii-art "Hello World"
```

## Testing

### Running Tests

```bash
# All tests
go test ./...

# With race detection
go test -race ./...

# Golden tests specifically
go test -run Golden ./...

# Coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Test Requirements

- Unit tests for all public functions
- Golden tests for output verification
- Coverage target: ≥80% overall, ≥90% for critical modules
- All tests must pass on Windows/Linux/macOS

## Code Style

### Required

- `gofmt` formatting (enforced in CI)
- `go vet` clean (no warnings)
- Follow Go naming conventions

### Recommended

- `staticcheck` clean
- `golangci-lint` clean
- Maximum function length: 50 lines
- Maximum file length: 500 lines

## Commit Guidelines

Use [Conventional Commits](https://conventionalcommits.org/):

```
feat: add shadow banner support
fix: handle empty string input correctly
docs: update API documentation
test: add golden tests for special characters
chore: update dependencies
```

## Pull Request Process

1. **Create feature branch**: `git checkout -b feature/your-feature`
2. **Small PRs**: Keep changes focused and reviewable
3. **Tests required**: All new code must have tests
4. **Documentation**: Update relevant docs
5. **CI must pass**: All checks must be green
6. **One reviewer**: Require at least one approval

### PR Checklist

- [ ] Tests added/updated and passing
- [ ] `gofmt` and `go vet` clean
- [ ] Documentation updated
- [ ] CHANGELOG.md updated (if user-facing)
- [ ] Golden tests updated (if output changed)
- [ ] Performance implications considered

## Issue Guidelines

### Bug Reports

Use the bug report template and include:

1. **Steps to reproduce**
2. **Input data** (exact string/file)
3. **Expected output**
4. **Actual output**
5. **Environment**: OS, Go version, ascii-art version

### Feature Requests

- Clear use case description
- Backward compatibility considerations
- Performance impact assessment

## Banner File Changes

Banner files are critical assets. For changes:

1. **Provenance**: Document source and licensing
2. **Format compliance**: Exactly 8 lines per character
3. **Testing**: Add comprehensive golden tests
4. **Validation**: Verify all printable ASCII chars (32-126)

## Release Process

1. Update CHANGELOG.md
2. Update version in code
3. Create PR with version bump
4. Tag release: `git tag v1.2.3`
5. CI builds and publishes artifacts
6. Update GitHub release notes

## Security

- Report security issues privately (see SECURITY.md)
- No credentials in code/tests
- Input validation for all user data
- Memory safety considerations

## Code Review Focus Areas

### For Reviewers

- **Security**: Input validation, memory safety
- **Performance**: Memory usage, CPU efficiency
- **Correctness**: Edge cases, error handling
- **Maintainability**: Code clarity, documentation
- **Testing**: Coverage, golden test accuracy

### Common Issues

- Missing error handling
- Inefficient string operations
- Inadequate input validation
- Missing edge case tests
- Outdated documentation

## Development Tools

### Recommended VS Code Extensions

- Go (official)
- Go Test Explorer
- Coverage Gutters

### Useful Commands

```bash
# Format all code
gofmt -w .

# Run linter
staticcheck ./...

# Generate test coverage
go test -coverprofile=coverage.out ./...

# Benchmark tests
go test -bench=. ./...
```

## Getting Help

- Open an issue for questions
- Check existing issues and PRs
- Review ARCHITECTURE.md for design decisions

## Recognition

Contributors get credit in:
- CHANGELOG.md for major contributions
- GitHub contributors page
- Release notes for new features

---

**Project maintained by Spiros Nikoloudakis (sp.nikoloudakis@gmail.com)**  
**Last updated: November 25, 2024**