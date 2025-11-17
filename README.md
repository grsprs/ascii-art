# Ascii-art

[![CI](https://github.com/username/ascii-art/workflows/CI/badge.svg)](https://github.com/username/ascii-art/actions)
[![Go Version](https://img.shields.io/badge/go-1.20+-blue.svg)](https://golang.org)
[![Coverage](https://img.shields.io/badge/coverage-85%25-green.svg)](https://github.com/username/ascii-art/actions)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Render text as ASCII banners using customizable banner files. Supports multi-line input, custom banners (standard/shadow/thinkertoy), and safe CLI usage.

## Quick Install

```bash
# Go install
go install github.com/username/ascii-art/cmd/ascii-art@latest

# Or download binary from releases
curl -L https://github.com/username/ascii-art/releases/latest/download/ascii-art-linux-amd64 -o ascii-art
chmod +x ascii-art
```

## Quick Usage

```bash
# Basic usage
ascii-art "Hello"

# Multi-line input
ascii-art "Hello\nWorld"

# Different banner style
ascii-art -banner=shadow "Hello"

# From file
ascii-art -file=input.txt
```

**Expected output for "Hello":**
```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```

## Banner Formats Supported

- **standard** - Classic ASCII art style
- **shadow** - Bold shadow effect
- **thinkertoy** - Decorative style

See [banners/](banners/) directory for format specifications.

## Features & Limits

- ✅ All printable ASCII characters (32-126)
- ✅ Multi-line input with `\n`
- ✅ Cross-platform (Windows/Linux/macOS)
- ✅ Input size limit: 1MB (configurable)
- ✅ Deterministic output
- ✅ Zero dependencies (stdlib only)

## CLI Reference

```bash
ascii-art [flags] [text]

Flags:
  -banner string    Banner style: standard, shadow, thinkertoy (default "standard")
  -file string      Read input from file
  -output string    Write output to file (default: stdout)
  -version          Show version information
  -help             Show help
```

## Exit Codes

- `0` - Success
- `1` - Invalid arguments
- `2` - I/O error
- `3` - Render error

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for development setup, testing, and contribution guidelines.

## Security

For security reports, see [SECURITY.md](SECURITY.md).

## License

MIT License - see [LICENSE](LICENSE) file.

## Maintainers

- [@maintainer](https://github.com/maintainer)

## Architecture

See [ARCHITECTURE.md](ARCHITECTURE.md) for technical details and design decisions.