# Go Package for Earthbucks Blockchain

## Overview

The `earthbucks-lib` package in Go is designed to implement the core functionality of the Earthbucks blockchain. This package is based on the existing TypeScript implementation, and our initial focus is on implementing data structures.

## Testing

To ensure the reliability and correctness of the package, we have integrated comprehensive testing.

### Testing Library

- **Testing Framework:** We use the Go testing framework, which is built into the Go standard library. This provides a straightforward way to write and run tests.

- **Additional Libraries:** We use the [testify](https://github.com/stretchr/testify) library for more expressive assertions and easier test management. This library helps in writing more readable and maintainable test cases.

  - **Library Installation:** To install testify, run:
    ```bash
    go get github.com/stretchr/testify
    ```

### Writing Tests

Tests are written in Go test files with the `_test.go` suffix. Each test file contains test functions that start with the `Test` prefix, following Go’s convention for test functions.

### Running Tests

To run all the tests in the project, use the following command:

```bash
go test ./...
```

- **Verbose Output:** For detailed output on test execution, use:
  ```bash
  go test -v ./...
  ```

- **Coverage Report:** To include code coverage information, run:
  ```bash
  go test -cover ./...
  ```

- **Timeout and Parallel Execution:** To set a timeout or control parallel test execution, use:
  ```bash
  go test -timeout 30s -parallel 4 ./...
  ```

### Why These Choices

- **Go Testing Framework:** The Go testing framework is a natural choice as it is integrated with the Go toolchain and supports a variety of testing needs out of the box.

- **Testify Library:** `testify` provides enhanced assertion methods, which help in writing more readable tests and simplifying test assertions. It also includes utilities for mocking and test suite management, which are beneficial for complex testing scenarios.

By using these tools and practices, we ensure that our implementation remains robust, maintainable, and aligned with industry standards.
