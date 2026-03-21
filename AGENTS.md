# AGENTS.md

This file provides guidelines and conventions for AI coding agents working with the ollama-image repository.

## Build, Lint, and Test Commands

### Taskfile
This project uses [Taskfile.yml](Taskfile.yml) for defining common commands.

**Build Commands:**
- `task build` - Go build without output
- `task install` - Install the application

**Test Commands:**
- `task test` - Run unit tests
- `task coverage` - Run unit tests with coverage
- `task coverage-html` - Generate test coverage as HTML
- `task open-coverage-html` - Generate coverage HTML and open it
- `task bench` - Run benchmarks on CPU and memory
- `task bench-profile` - Run benchmarks with memory and CPU profiling

**Lint and Security:**
- `task lint` - Run golangci-lint
- `task sec` - Security check with gosec

### Go Commands
- `go build ./...` - Build all packages
- `go test ./...` - Run all tests
- `go test -cover ./...` - Run all tests with coverage
- `go test -bench=. -benchmem ./...` - Run benchmarks with memory profiling
- `golangci-lint run` - Run linter
- `gosec ./...` - Security check

### Running a Single Test
To run a specific test function, use:
```bash
go test -v -run TestFunctionName
```

To run tests in a specific package:
```bash
go test -v ./path/to/package
```

## Code Style Guidelines

### Go Conventions
- **Formatting**: Use `gofmt` for code formatting
- **Imports**: Group imports by standard library, external packages, and internal packages
- **Naming**: Use camelCase for variables and functions, PascalCase for types
- **Error Handling**: Return errors explicitly and check them
- **Comments**: Use clear, concise comments for imports, functions, and types

### Import Grouping
```go
// Standard library
import (
	"fmt"
	"os"
)

// External packages
import (
	"github.com/spf13/cobra"
	"github.com/ollama/ollama/api"
)

// Internal packages
import (
	"github.com/alexhokl/helper/cli"
)
```

### Types and Naming
- Use descriptive names for types and variables
- Prefer specific types over empty interfaces
- Use pointer receivers for methods that modify state
- Export only necessary types and functions

### Error Handling
- Wrap errors with context using fmt.Errorf
- Handle errors explicitly rather than ignoring them
- Use sentinel errors for common error cases
- Log errors appropriately with logging levels

### Test Files
- Name test files with _test suffix: `file_test.go`
- Use Table-driven tests for multiple test cases
- Test both happy path and error cases
- Mock external dependencies when necessary

## Project Structure

```
/ollama-image
├── cmd/
│   ├── ask.go
│   ├── describe.go
│   ├── root.go
│   └── ...
├── main.go
├── Taskfile.yml
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

### Package Structure
- `cmd/`: Contains command-line interface code using Cobra
- `main.go`: Entry point of the application
- Taskfile.yml: Task automation configuration

## Development Workflow

1. **Writing Code**:
   - Follow Go idioms and best practices
   - Use meaningful variable and function names
   - Write clear, concise code with appropriate comments
   - Group related functionality together

2. **Testing**:
   - Write tests for all new functionality
   - Test both success and error cases
   - Use table-driven tests for multiple scenarios
   - Ensure tests are deterministic and fast

3. **Review**:
   - Run linters before committing
   - Check that code follows project conventions
   - Ensure all tests pass
   - Verify error handling is robust

## Dependencies
- The project uses Go modules for dependency management
- Add new dependencies with `go get` command
- Keep dependencies updated regularly
- Verify third-party dependencies are from trusted sources

## Security Considerations
- Validate all inputs
- Sanitize outputs before displaying to users
- Use secure defaults for all configuration
- Regularly run security checks (gosec)

## Command Completion
Available commands for completion:
- `task completion-mac` - Generate bash completion for macOS
- `task completion-linux` - Generate bash completion for Linux
