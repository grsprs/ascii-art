# Architecture Documentation

## High-Level Overview

```
Input → Parser → Renderer → Output Writer
  ↓       ↓         ↓           ↓
 CLI    Banner    ASCII      stdout/file
Args    Loader    Builder
```

## System Components

### 1. CLI Interface (`cmd/ascii-art/`)
- Argument parsing and validation
- Input source selection (args, stdin, file)
- Output destination handling
- Error reporting and exit codes

### 2. Banner Management (`internal/banner/`)
- Banner file loading and parsing
- Character glyph mapping (rune → 8-line ASCII)
- Banner format validation
- Caching for performance

### 3. Renderer (`internal/renderer/`)
- Text-to-ASCII conversion logic
- Multi-line text handling
- Character positioning and spacing
- Output buffer management

### 4. Validation (`internal/validation/`)
- Input size limits (default: 1MB)
- Character range validation (ASCII 32-126)
- Banner file integrity checks
- Security input sanitization

## Data Models

### Glyph Structure
```go
type Glyph struct {
    Lines [8]string  // Exactly 8 lines per character
    Width int        // Character width in columns
}
```

### Banner Structure
```go
type Banner struct {
    Name   string
    Glyphs map[rune]Glyph  // ASCII 32-126 mapping
}
```

## Banner File Format Specification

### Format Rules
- **Height**: Exactly 8 lines per character
- **Separation**: Characters separated by single newline
- **Encoding**: UTF-8
- **Range**: ASCII characters 32-126 (95 total)
- **Order**: Sequential from space (32) to tilde (126)

### Example Format
```
      
      
      
      
      
      
      
      

 _  
| | 
| | 
| | 
|_| 
(_) 
    
    
```

### Missing Character Handling
- Undefined characters render as space equivalent
- Banner files must include all printable ASCII
- Validation ensures completeness

## Processing Flow

### 1. Input Processing
```
Text Input → Character Validation → Size Check → Parse Lines
```

### 2. Banner Loading
```
Banner File → Format Validation → Parse Glyphs → Cache in Memory
```

### 3. Rendering
```
For each line of input:
  For each character in line:
    Lookup glyph in banner
    Append to output buffer
  Add newline separation
```

### 4. Output Generation
```
Buffer → Line Assembly → Final Output → Write to Destination
```

## Memory Management

### Design Decisions
- **In-memory processing**: Full input loaded for simplicity
- **Banner caching**: Loaded once, reused for multiple renders
- **Streaming output**: Results written incrementally
- **Memory limits**: 1MB input cap prevents OOM

### Performance Characteristics
- **Time Complexity**: O(n) where n = input length
- **Space Complexity**: O(n × 8) for output buffer
- **Banner loading**: O(1) amortized with caching

## Error Handling

### Error Categories
| Exit Code | Category | Description |
|-----------|----------|-------------|
| 0 | Success | Normal completion |
| 1 | Invalid Args | CLI argument errors |
| 2 | I/O Error | File read/write failures |
| 3 | Render Error | Banner/rendering issues |

### Error Propagation
- Errors bubble up through call stack
- Context preserved with error wrapping
- User-friendly messages at CLI level
- Detailed logging for debugging

## Security Architecture

### Input Validation
- Size limits enforced at entry points
- Character range validation (32-126)
- Path traversal prevention for file inputs
- No code execution in banner files

### Resource Protection
- Memory usage monitoring
- CPU time limits (future enhancement)
- File descriptor limits
- Temporary file cleanup

## Thread Safety

### Current State
- **Single-threaded**: No concurrency in v1.0
- **Banner loading**: Thread-safe read-only after init
- **No shared state**: Each invocation independent

### Future Considerations
- Concurrent banner loading
- Parallel line rendering
- Thread-safe caching

## Dependencies

### Standard Library Only
- `os` - File operations and CLI args
- `fmt` - String formatting and output
- `strings` - String manipulation
- `bufio` - Buffered I/O operations
- `path/filepath` - Safe path handling

### Rationale
- **Audit simplicity**: No external dependencies
- **Security**: Reduced attack surface
- **Maintenance**: No dependency updates needed
- **Portability**: Works everywhere Go works

## Performance Considerations

### Optimization Strategies
- Banner file caching
- Efficient string building
- Minimal memory allocations
- Early validation to fail fast

### Benchmarks (Target)
- Small input (<100 chars): <1ms
- Medium input (1KB): <10ms
- Large input (1MB): <100ms
- Memory usage: <10MB peak

## Extensibility Points

### Future Enhancements
- Custom banner file formats
- Color output support
- Font size scaling
- Output format options (HTML, SVG)

### Plugin Architecture (Future)
- Banner format plugins
- Output format plugins
- Input source plugins
- Validation rule plugins

## Testing Strategy

### Unit Tests
- Each component tested in isolation
- Mock dependencies for pure unit tests
- Edge cases and error conditions
- Performance regression tests

### Integration Tests
- End-to-end CLI testing
- Golden file comparisons
- Cross-platform compatibility
- Real banner file testing

### Property Tests
- Input/output relationship validation
- Deterministic output verification
- Memory usage bounds checking
- Performance characteristic validation

## Deployment Architecture

### Build Process
```
Source Code → Go Build → Static Binary → Multi-platform Release
```

### Distribution
- GitHub Releases with binaries
- Go module for library usage
- Container images (future)
- Package managers (future)

### Configuration
- CLI flags for runtime options
- Environment variables for defaults
- Configuration files (future enhancement)