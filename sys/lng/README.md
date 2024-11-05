# `lng` - Language Processing Component

## Overview
A sophisticated language processing system implemented in Go, featuring lexical analysis, parsing, and compilation capabilities for multiple domain-specific languages (DSLs). This component is part of a larger algorithmic trading engine and handles configuration, JSON processing, and a custom programming language.

## Key Features

### Scanner System (`scn`)
- UTF-8 compliant lexical analyzer with full Unicode support
- Line and column tracking for precise error reporting
- Built-in handling of byte order marks (BOM)
- Efficient escape sequence processing
- Advanced token boundary detection

### Parser Framework
Implements parsers for multiple formats:

#### Custom Programming Language (`pro`)
- Complete lexical analysis and parsing pipeline
- Expression parsing and evaluation
- Scope management and symbol resolution
- Type system with support for:
  - Basic types (string, boolean, float, int)
  - Complex types (arrays, objects)
  - Domain-specific types (time ranges, bounds)

#### JSON Parser (`jsn`)
- Full JSON specification compliance
- Recursive descent parser with look-ahead
- Nested object and array support
- Type-safe value extraction
- Path-based value lookup system

#### Configuration Parser (`cfg`)
- Structured configuration file processing
- Hierarchical key-value parsing
- Support for multiple data types:
  - Strings, Booleans, Floats
  - Integers, Unsigned integers
  - Time values and ranges
  - Arrays and nested objects

### Type System
- Strong static typing
- Generic type support
- Extensible type definitions
- Built-in type validation

### Technical Specifications
- Zero external dependencies for core parsing
- Comprehensive test coverage
- Memory-efficient design

## Implementation Details

### Scanner Architecture
```go
type Scn struct {
    Ch   rune  // Current character
    Size int   // Size in bytes of current character
    Idx  unt.Unt  // Current index in input
    Ln   unt.Unt  // Current line number
    Col  unt.Unt  // Current column number
    End  bool  // End of input indicator
}
```

### Key Components
1. **Tokenizer**: Converts raw input into token streams
2. **Parser**: Processes token streams into structured data
3. **Type System**: Manages type checking and validation
4. **Scope Manager**: Handles symbol tables and scope resolution

## Performance Features
- Efficient UTF-8 decoding
- Optimized string handling
- Minimal memory allocations
- Fast lookup algorithms for nested structures

## Usage Examples

### JSON Parsing
```go
var jsnr jsn.Jsnr
jsnr.Reset(jsonString)
value := jsnr.Str("path", "to", "value")
```

### Configuration Processing
```go
var cfgr cfg.Cfgr
cfgr.Reset(configString)
setting := cfgr.Flt("section", "subsection", "parameter")
```

## Engineering Highlights
- Clean separation of concerns between scanning, parsing, and evaluation
- Extensive error handling with detailed diagnostics
- Modular design enabling easy extension
- Comprehensive unit testing
- Production-grade implementation used in algorithmic trading

## Technical Competencies Demonstrated
- Advanced Go programming
- Compiler design and implementation
- Type system design
- Performance optimization
- Large-scale system architecture
- Test-driven development
- Domain-specific language design

This component showcases strong software engineering principles while implementing complex language processing capabilities, making it a significant demonstration of technical expertise in Go development and compiler construction.