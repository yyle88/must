[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/must/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/must/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/must)](https://pkg.go.dev/github.com/yyle88/must)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/must/master.svg)](https://coveralls.io/github/yyle88/must?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/must.svg)](https://github.com/yyle88/must/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/must)](https://goreportcard.com/report/github.com/yyle88/must)

# `must` – A Simple Assertion Library for Go

The must library simplifies assertions and panics on failure. Assert conditions directly, avoiding long checks

---

## CHINESE README

[中文说明](README.zh.md)

## Features

- **Simple Error Handling**: Assert conditions directly, avoiding long checks.
- **Quick Feedback**: Catch bugs early with clear panic messages.
- **Lightweight & Fast**: Minimal code overhead for speed. Very easy to use.
- **Versatile Assertions**: Supports checking non-zero values, slice lengths.

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

	println("Array length is correct")
}
```

---

## Core Assertions

Here are the core assertions in `must`, summarized in a table:

| **Function**                 | **Description**                                            | **Example**                   | **Notes**                              |
|------------------------------|------------------------------------------------------------|-------------------------------|----------------------------------------|
| **`True(v bool)`**           | Panics if `v` is false.                                    | `must.True(isValid)`          | Validates if `v` is `true`.            |
| **`Done(err error)`**        | Panics if `err` is not nil.                                | `must.Done(err)`              | Ensures no error occurred.             |
| **`Must(err error)`**        | Panics if `err` is not nil.                                | `must.Must(err)`              | Similar to `Done`.                     |
| **`Nice(a V)`**              | Panics if `a` is zero.                                     | `must.Nice(value)`            | Ensures `a` is non-zero.               |
| **`Zero(a V)`**              | Panics if `a` is not zero.                                 | `must.Zero(value)`            | Ensures `a` is zero.                   |
| **`None(a V)`**              | Panics if `a` is non-zero.                                 | `must.None(value)`            | Ensures `a` is zero.                   |
| **`Null(v any)`**            | Panics if `v` is not `nil`.                                | `must.Null(ptr)`              | Ensures `v` is `nil`.                  |
| **`Full(v any)`**            | Panics if `v` is `nil`.                                    | `must.Full(value)`            | Ensures `v` is non-`nil`.              |
| **`Equals(a, b V)`**         | Panics if `a` and `b` are not equal.                       | `must.Equals(a, b)`           | Checks if `a` equals `b`.              |
| **`Same(a, b V)`**           | Panics if `a` and `b` are not equal.                       | `must.Same(a, b)`             | Alias of `Equals`.                     |
| **`Is(a, b V)`**             | Panics if `a` and `b` are not equal.                       | `must.Is(a, b)`               | Alias of `Equals`.                     |
| **`Ise(err, target error)`** | Panics if `err` does not match `target` using `errors.Is`. | `must.Ise(err, targetErr)`    | Error matching similar to `errors.Is`. |
| **`Ok(a V)`**                | Panics if `a` is zero.                                     | `must.Ok(value)`              | Ensures `a` is non-zero.               |
| **`OK(a V)`**                | Alias of `Ok`, checks for non-zero value.                  | `must.OK(value)`              | Similar to `Ok`.                       |
| **`TRUE(v bool)`**           | Panics if `v` is false.                                    | `must.TRUE(isValid)`          | Alias of `True`.                       |
| **`FALSE(v bool)`**          | Panics if `v` is true.                                     | `must.FALSE(isError)`         | Ensures `v` is `false`.                |
| **`False(v bool)`**          | Panics if `v` is true.                                     | `must.False(isError)`         | Similar to `FALSE`.                    |
| **`Have(a []T)`**            | Panics if `a` is empty.                                    | `must.Have(slice)`            | Ensures `a` is not empty.              |
| **`Length(a []T, n int)`**   | Panics if `a` length is not `n`.                           | `must.Length(slice, 3)`       | Ensures `a` length is `n`.             |
| **`Len(a []T, n int)`**      | Alias of `Length`, ensures `a` length is `n`.              | `must.Len(slice, 3)`          | Validates `a` length.                  |
| **`In(v T, a []T)`**         | Panics if `v` is not in `a`.                               | `must.In(value, slice)`       | Ensures `v` is in `a`.                 |
| **`Contains(a []T, v T)`**   | Panics if `a` does not contain `v`.                        | `must.Contains(slice, value)` | Ensures `a` contains `v`.              |

---

## License

MIT License. See [LICENSE](LICENSE).

---

## Contribute and Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with `must`!** 🎉
