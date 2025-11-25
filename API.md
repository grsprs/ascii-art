# CLI Reference

Complete guide to using the ascii-art command line tool.

## Command Syntax

```bash
ascii-art [flags] [text]
```

## Flags

### `-banner string`
**Description**: Specify banner style  
**Default**: `standard`  
**Valid values**: `standard`, `shadow`, `thinkertoy`  
**Example**: `ascii-art -banner=shadow "Hello"`

### `-file string`
**Description**: Read input from file instead of arguments  
**Default**: None (use command line arguments)  
**Example**: `ascii-art -file=input.txt`

### `-output string`
**Description**: Write output to file instead of stdout  
**Default**: stdout  
**Example**: `ascii-art -output=result.txt "Hello"`

### `-version`
**Description**: Display version information and exit  
**Example**: `ascii-art -version`

### `-help`
**Description**: Display help information and exit  
**Example**: `ascii-art -help`

## Input Sources

### Priority Order
1. `-file` flag (highest priority)
2. Command line arguments
3. stdin (if no args and no file)

### Input Validation
- **Size limit**: 1MB maximum
- **Character range**: ASCII 32-126 (printable characters)
- **Newlines**: `\n` supported for multi-line output
- **Empty input**: Produces empty output (exit code 0)

## Output Destinations

### stdout (default)
```bash
ascii-art "Hello" > output.txt
```

### File output
```bash
ascii-art -output=result.txt "Hello"
```

## Exit Codes

| Code | Meaning | Description |
|------|---------|-------------|
| 0 | Success | Normal completion |
| 1 | Invalid Arguments | CLI usage errors |
| 2 | I/O Error | File read/write failures |
| 3 | Render Error | Banner or rendering issues |

## Usage Examples

### Basic Usage
```bash
# Simple text
ascii-art "Hello"

# Multi-line text
ascii-art "Hello\nWorld"

# Empty string (valid, produces no output)
ascii-art ""
```

### Banner Styles
```bash
# Standard style (default)
ascii-art "Hello"
ascii-art -banner=standard "Hello"

# Shadow style
ascii-art -banner=shadow "Hello"

# Thinkertoy style
ascii-art -banner=thinkertoy "Hello"
```

### File Operations
```bash
# Read from file
echo "Hello World" > input.txt
ascii-art -file=input.txt

# Write to file
ascii-art -output=banner.txt "Hello"

# Both file input and output
ascii-art -file=input.txt -output=banner.txt
```

### Complex Examples
```bash
# Numbers and special characters
ascii-art "Hello 123!"

# Mixed case
ascii-art "HeLLo WoRLd"

# Special characters
ascii-art "Hello, World! @#$%"

# Brackets and symbols
ascii-art "{Hello [World]}"
```

## Input Format Specifications

### Supported Characters
- **Range**: ASCII 32-126
- **Count**: 95 printable characters
- **Special**: Space, punctuation, numbers, letters

### Character Map
```
 !"#$%&'()*+,-./0123456789:;<=>?
@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_
`abcdefghijklmnopqrstuvwxyz{|}~
```

### Escape Sequences
- `\n` - Newline (creates separate banner lines)
- `\\` - Literal backslash
- `\"` - Literal quote (in shell)

### Multi-line Behavior
```bash
# Input with \n
ascii-art "Line1\nLine2"

# Produces:
# [Line1 rendered in ASCII]
# [8 blank lines for spacing]
# [Line2 rendered in ASCII]
```

## Error Handling

### Invalid Arguments
```bash
# Unknown flag
ascii-art -invalid "Hello"
# Exit code: 1, Error message displayed

# Invalid banner
ascii-art -banner=invalid "Hello"  
# Exit code: 1, Error: unknown banner style
```

### I/O Errors
```bash
# File not found
ascii-art -file=missing.txt
# Exit code: 2, Error: file not found

# Permission denied
ascii-art -output=/root/output.txt "Hello"
# Exit code: 2, Error: permission denied
```

### Render Errors
```bash
# Corrupted banner file
ascii-art -banner=corrupted "Hello"
# Exit code: 3, Error: invalid banner format
```

## Performance Specifications

### Input Limits
- **Maximum size**: 1MB (1,048,576 bytes)
- **Recommended**: <10KB for optimal performance
- **Memory usage**: ~10x input size for output buffer

### Response Times (Target)
- Small input (<100 chars): <1ms
- Medium input (1KB): <10ms  
- Large input (100KB): <100ms
- Maximum input (1MB): <1s

## Compatibility

### Operating Systems
- Linux (x86_64, ARM64)
- macOS (x86_64, ARM64)
- Windows (x86_64)

### Go Versions
- Minimum: Go 1.20
- Tested: Go 1.20, 1.21, 1.22
- Recommended: Latest stable

## Environment Variables

### Configuration
```bash
# Default banner style
export ASCII_ART_BANNER=shadow
ascii-art "Hello"  # Uses shadow banner

# Input size limit (bytes)
export ASCII_ART_MAX_SIZE=2097152  # 2MB
```

## Integration Examples

### Shell Scripts
```bash
#!/bin/bash
# Generate banner for script output
ascii-art "$(basename "$0")" 
echo "Script starting..."
```

### Makefiles
```makefile
banner:
	@ascii-art "Build Complete"
```

### CI/CD Pipelines
```yaml
- name: Success Banner
  run: ascii-art "Build Success!"
```

## API Stability

### Versioning
- Semantic versioning (semver)
- CLI interface stability guaranteed
- Flag compatibility maintained
- Output format consistency

### Deprecation Policy
- 6-month notice for breaking changes
- Migration guides provided
- Legacy flag support during transition

---

**CLI designed by Spiros Nikoloudakis (sp.nikoloudakis@gmail.com)**  
**Last updated: November 25, 2024**