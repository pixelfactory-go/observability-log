# Contributing to Observability Log

Thank you for considering contributing to Observability Log! This document provides guidelines and instructions for contributing to this project.

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for all contributors. Please be professional, considerate, and constructive in all interactions.

## How to Contribute

### Reporting Issues

If you find a bug or have a feature request:

1. **Search existing issues** to avoid duplicates
2. **Create a new issue** with a clear title and description
3. **Provide details**:
   - For bugs: steps to reproduce, expected vs actual behavior, environment details
   - For features: use case, proposed solution, potential alternatives
4. **Add labels** if you have permission (bug, enhancement, documentation, etc.)

### Submitting Pull Requests

1. **Fork the repository** and create a new branch from `main`
2. **Make your changes** following the coding standards below
3. **Test your changes** thoroughly
4. **Update documentation** if needed
5. **Write a clear commit message** following our commit guidelines
6. **Submit a pull request** with a detailed description

#### Pull Request Guidelines

- Keep PRs focused on a single feature or fix
- Reference related issues in the PR description (e.g., "Fixes #123")
- Ensure all tests pass and code is properly formatted
- Add tests for new functionality
- Update relevant documentation
- Wait for review and address feedback promptly

## Development Setup

### Prerequisites

- **Go 1.23 or higher**: [Download](https://go.dev/dl/)
- **golangci-lint v2.8.0 or higher**: [Installation guide](https://golangci-lint.run/docs/welcome/install/)
- **Make**: Usually pre-installed on Unix systems

### Local Development

1. **Clone the repository**:
   ```bash
   git clone https://github.com/pixelfactory-go/observability-log.git
   cd observability-log
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Run tests**:
   ```bash
   make test
   ```

4. **Format code**:
   ```bash
   make fmt
   ```

5. **Run linter**:
   ```bash
   make lint
   ```

### Available Make Targets

- `make fmt` - Format code using gofmt
- `make lint` - Run golangci-lint
- `make test` - Run tests with coverage
- `make build` - Build the project
- `make help` - Show all available targets

## Coding Standards

### Go Guidelines

- Follow the [Effective Go](https://go.dev/doc/effective_go) guidelines
- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` for formatting (run `make fmt`)
- Pass all linter checks (run `make lint`)

### Code Style

- **Package names**: lowercase, single word when possible
- **Variable names**: camelCase for local variables, PascalCase for exported identifiers
- **Error handling**: always check errors, don't ignore them
- **Comments**:
  - All exported functions, types, and constants must have godoc comments
  - Comments should be complete sentences ending with periods
  - Start comments with the name of the thing being described

### Documentation

- Document all exported functions, types, and constants
- Use examples in godoc when helpful
- Keep documentation up to date with code changes
- Write clear, concise explanations

### Testing

- Write unit tests for all new functionality
- Maintain or improve code coverage
- Use table-driven tests where appropriate
- Test both success and failure cases
- Test edge cases and boundary conditions

Example test structure:
```go
func TestFunctionName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "valid input",
			input:   "test",
			want:    "expected",
			wantErr: false,
		},
		// Add more test cases...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FunctionName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FunctionName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FunctionName() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

## Commit Message Guidelines

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification for commit messages.

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `chore`: Maintenance tasks, dependency updates
- `ci`: CI/CD changes
- `refactor`: Code refactoring without feature changes
- `test`: Adding or updating tests
- `perf`: Performance improvements

### Examples

```
feat(fields): add support for custom ECS fields

Implement a new Field() function that allows users to add
arbitrary ECS-compliant fields to their log entries.

Closes #42
```

```
fix: correct nil pointer dereference in HTTPRequest

The HTTPRequest field helper was not checking for nil URL
before accessing its properties, causing a panic.

Fixes #156
```

```
docs: update README with new field helpers

Add documentation for recently added field helper functions
and improve existing examples.
```

### Commit Message Rules

- Use the imperative mood ("add feature" not "added feature")
- Keep the subject line under 50 characters
- Capitalize the subject line
- Don't end the subject line with a period
- Separate subject from body with a blank line
- Wrap the body at 72 characters
- Use the body to explain what and why, not how

## Review Process

### For Contributors

1. Submit your PR and wait for initial review
2. Address reviewer feedback promptly
3. Request re-review after making changes
4. Be patient and respectful of reviewers' time

### For Reviewers

1. Review PRs promptly (within 2-3 business days)
2. Provide constructive, specific feedback
3. Approve when ready or request changes clearly
4. Re-review promptly after changes are made

### Merge Requirements

- All tests must pass
- All linter checks must pass
- At least one approval from a maintainer
- All conversations resolved
- Up to date with main branch

## Questions?

If you have questions that aren't covered here:

1. Check the [README](README.md) for basic information
2. Search [existing issues](https://github.com/pixelfactory-go/observability-log/issues)
3. Create a new issue with your question

## License

By contributing to this project, you agree that your contributions will be licensed under the MIT License.
