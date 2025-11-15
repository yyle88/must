[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/must/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/must/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/must)](https://pkg.go.dev/github.com/yyle88/must)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/must/main.svg)](https://coveralls.io/github/yyle88/must?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23%2C%201.24%2C%201.25-lightgrey.svg)](https://github.com/yyle88/must)
[![GitHub Release](https://img.shields.io/github/release/yyle88/must.svg)](https://github.com/yyle88/must/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/must)](https://goreportcard.com/report/github.com/yyle88/must)

# must

Simple assertion utilities with panic-on-failure semantics designed to reduce boilerplate code in Go.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Main Features

🎯 **Panic-on-Failure Validation**: Clean assertions with automatic panic on failure when conditions are violated
⚡ **Type-Safe Generics**: Comprehensive support with Go generics spanning assertion types
🔄 **Stack Frame Adjustment**: Precise panic location through intelligent skip configurations
🌍 **Structured Logging**: Deep integration with zap providing detailed panic context
📋 **Domain-Specific Packages**: Specialized utilities with numeric, string, slice, and map operations

---

## Installation

```bash
go get github.com/yyle88/must
```

---

## Quick Start

### Example 1: Assert a Non-Zero Value

```go
package main

import (
	"github.com/yyle88/must"
)

func main() {
	value := 42
	must.Nice(value) // Panics if value is zero

	println("Value is valid:", value)
}
```

---

### Example 2: Validate No Error

```go
package main

import (
	"errors"
	"github.com/yyle88/must"
)

func main() {
	err := someFunction()
	must.Done(err) // Panics if err is not nil

	println("No error encountered!")
}

func someFunction() error {
	return errors.New("unexpected error")
}
```

---

### Example 3: Check Slice Length

```go
package main

import (
	"github.com/yyle88/must"
)

func main() {
	arr := []int{1, 2, 3}
	must.Length(arr, 3) // Panics if the length is not 3

	println("Slice length is correct")
}
```

---

## Core Assertions

Here are the core assertions in `must`, summarized in a table:

| **Function**                 | **Description**                                            | **Example**                   | **Notes**                              |
|------------------------------|------------------------------------------------------------|-------------------------------|----------------------------------------|
| **`True(v bool)`**           | Panics if `v` is false.                                    | `must.True(isValid)`          | Validates if `v` is `true`.            |
| **`Done(err error)`**        | Panics if `err` is not nil.                                | `must.Done(err)`              | Ensures no error occurred.             |
| **`Must(err error)`**        | Panics if `err` is not nil.                                | `must.Must(err)`              | Same as `Done`.                        |
| **`Nice(a V)`**              | Panics if `a` is zero.                                     | `must.Nice(value)`            | Ensures `a` is non-zero.               |
| **`Zero(a V)`**              | Panics if `a` is not zero.                                 | `must.Zero(value)`            | Ensures `a` is zero.                   |
| **`None(a V)`**              | Panics if `a` is non-zero.                                 | `must.None(value)`            | Ensures `a` is zero.                   |
| **`Null(v any)`**            | Panics if `v` is not `nil`.                                | `must.Null(ptr)`              | Ensures `v` is `nil`.                  |
| **`Full(v any)`**            | Panics if `v` is `nil`.                                    | `must.Full(value)`            | Ensures `v` is non-`nil`.              |
| **`Equals(a, b V)`**         | Panics if `a` and `b` are not the same.                    | `must.Equals(a, b)`           | Checks if `a` equals `b`.              |
| **`Same(a, b V)`**           | Panics if `a` and `b` are not the same.                    | `must.Same(a, b)`             | Alias of `Equals`.                     |
| **`SameNice(a, b V)`**       | Panics if `a` and `b` are not the same, both non-zero.     | `must.SameNice(a, b)`         | Ensures same and non-zero.             |
| **`Sane(a, b V)`**           | Panics if `a` and `b` are not the same, both non-zero.     | `must.Sane(a, b)`             | Alias of `SameNice`.                   |
| **`Diff(a, b V)`**           | Panics if `a` and `b` are the same.                        | `must.Diff(a, b)`             | Ensures values mismatch.               |
| **`Different(a, b V)`**      | Panics if `a` and `b` are the same.                        | `must.Different(a, b)`        | Alias of `Diff`.                       |
| **`Is(a, b V)`**             | Panics if `a` and `b` are not the same.                    | `must.Is(a, b)`               | Alias of `Equals`.                     |
| **`Ise(err, target error)`** | Panics if `err` does not match `target` using `errors.Is`. | `must.Ise(err, targetErr)`    | Matching like `errors.Is` function.    |
| **`Ok(a V)`**                | Panics if `a` is zero.                                     | `must.Ok(value)`              | Ensures `a` is non-zero.               |
| **`OK(a V)`**                | Alias of `Ok`, checks non-zero value.                      | `must.OK(value)`              | Same as `Ok`.                          |
| **`TRUE(v bool)`**           | Panics if `v` is false.                                    | `must.TRUE(isValid)`          | Alias of `True`.                       |
| **`FALSE(v bool)`**          | Panics if `v` is true.                                     | `must.FALSE(isError)`         | Ensures `v` is `false`.                |
| **`False(v bool)`**          | Panics if `v` is true.                                     | `must.False(isError)`         | Same as `FALSE`.                       |
| **`Have(a []T)`**            | Panics if `a` has no elements.                             | `must.Have(slice)`            | Ensures `a` is not vacant.             |
| **`Length(a []T, n int)`**   | Panics if `a` length is not `n`.                           | `must.Length(slice, 3)`       | Ensures `a` length is `n`.             |
| **`Len(a []T, n int)`**      | Alias of `Length`, ensures `a` length is `n`.              | `must.Len(slice, 3)`          | Validates `a` length.                  |
| **`In(v T, a []T)`**         | Panics if `v` is not in `a`.                               | `must.In(value, slice)`       | Ensures `v` is in `a`.                 |
| **`Contains(a []T, v T)`**   | Panics if `a` does not contain `v`.                        | `must.Contains(slice, value)` | Ensures `a` contains `v`.              |

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a mistake?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yyle88/must.svg?variant=adaptive)](https://starchart.cc/yyle88/must)
