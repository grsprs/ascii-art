# ASCII Art Generator

Convert text into ASCII art banners with different font styles.

## Installation

```bash
go install github.com/grsprs/ascii-art/cmd/ascii-art@latest
```

## Usage

```bash
ascii-art "Hello"
ascii-art -banner=shadow "Hello"
ascii-art -file=input.txt
ascii-art -output=output.txt "Hello"
```

## Options

- `-banner` - Font style: standard, shadow, thinkertoy (default: standard)
- `-file` - Read input from file
- `-output` - Write output to file
- `-version` - Show version
- `-help` - Show help

## Example

```bash
ascii-art "Hello"
```

Output:
```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```

## Features

- All printable ASCII characters (32-126)
- Multi-line input with `\n`
- Three font styles
- Cross-platform compatibility
- Zero dependencies
- 1MB input limit
- Secure path handling

## Development

```bash
git clone https://github.com/grsprs/ascii-art.git
cd ascii-art
go test ./...
go run ./cmd/ascii-art "Test"
```

## Building

```bash
go build -o ascii-art ./cmd/ascii-art
```

## License

MIT License - see LICENSE file.

## Author

Spiros Nikoloudakis (sp.nikoloudakis@gmail.com)