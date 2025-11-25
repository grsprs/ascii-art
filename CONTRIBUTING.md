# Contributing

## Development Setup

```bash
git clone https://github.com/grsprs/ascii-art.git
cd ascii-art
go test ./...
go run ./cmd/ascii-art "Hello"
```

## Code Standards

- Use `gofmt` for formatting
- Pass `go vet` checks
- Follow Go naming conventions
- Add tests for new features

## Pull Request Process

1. Create feature branch
2. Add tests
3. Ensure all tests pass
4. Submit pull request

## Testing

```bash
go test ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Building

```bash
go build -o ascii-art ./cmd/ascii-art
```