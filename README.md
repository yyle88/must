# `must` – A Simple, Yet Powerful Assertion Library for Go

### Why settle for less when you can be sure?

In Go, errors are part of the game, but sometimes you want to enforce guarantees that your code must meet—without the need for extensive checks or verbose error-handling. That's where `must` comes in: a lightweight package that ensures certain conditions *must* be true for your program to continue, panic-free. If the condition fails, you get an immediate and clear failure, saving time debugging or testing edge cases.

## README
[中文说明](README.zh.md)

## Install

Quickly integrate `must` into your Go project by running:

```bash
go get github.com/yyle88/must
```

## Core Philosophy

The premise behind `must` is simple: **assert the unassertable**.

In many cases, Go developers use error handling to validate that conditions are met, but what if you don’t want to write multiple checks for simple conditions? What if certain things in your codebase *must* happen, and anything less should halt the program?

That’s exactly what `must` does. It brings you immediate feedback if something goes wrong with minimal boilerplate, ensuring you catch bugs early in the development cycle.

## Features

- **Enforce Certain Conditions**: Assert that a value or state must meet specific criteria (e.g., non-nil, non-zero).
- **Clear Failures**: Panic when an assertion fails, providing an explicit and actionable message.
- **Optimized for Speed**: Perfect for prototypes and small applications where fast feedback is critical.

## Usage

Here’s how you can use `must` in your code to handle assertions concisely:

### Example 1: Validate Non-Nil Values

```go
package main

import (
	"fmt"
	"github.com/yyle88/must"
)

func main() {
	// Simulate a function that may return an error
	result, err := someFunction()
	must.Done(err)   // Panics if err is non-nil
	must.Nice(result) // Panics if result is the zero value (nil or empty)

	fmt.Println("Result is:", result)
}

func someFunction() (string, error) {
	// For demonstration, returning a result and no error
	return "Hello, must!", nil
}
```

### Example 2: Assert a Boolean Condition

```go
package main

import (
	"fmt"
	"github.com/yyle88/must"
)

func main() {
	// Simulate a check that should return true
	isActive := checkStatus()
	must.True(isActive)  // Panics if isActive is false

	fmt.Println("Status is active:", isActive)
}

func checkStatus() bool {
	// Returning true for demonstration
	return true
}
```

### Example 3: Ensure Length Matches

```go
package main

import (
	"fmt"
	"github.com/yyle88/must"
)

func main() {
	// Simulating an array
	arr := []int{1, 2, 3}
	must.Length(arr, 3) // Panics if the length is not 3

	fmt.Println("Array length is correct:", len(arr))
}
```

## Why `must`?

- **Simplicity Over Verbosity**: In Go, error handling is often verbose. `must` simplifies assertions into short, easy-to-understand statements.
- **Immediate Feedback**: Perfect for scenarios where you need to catch conditions early—be it during development or in fast-paced prototypes.
- **Zero-Cost Abstraction**: `must` introduces no additional complexity and is highly idiomatic, keeping your Go codebase clean and easy to maintain.

## Core Assertions

- **`must.True(condition bool)`**: Asserts that a boolean condition is `true`. Panics if `false`.
- **`must.Done(err error)`**: Verifies that an error is `nil`. Panics if the error is not `nil`.
- **`must.Nice(value interface{})`**: Asserts that the value is not the zero value (i.e., empty or uninitialized). Panics if zero.
- **`must.Length(array interface{}, length int)`**: Ensures the provided array or slice has the expected length.
- **`must.Equals(a, b interface{})`**: Compares two values for equality and panics if they are not equal.

## A Word of Caution

While `must` is great for quick feedback during development or testing, **it should not be used for production error handling**. If you need fine-grained control over errors, consider using traditional Go error-handling techniques.

## Future Plans

- Expand assertion types, like checking for non-empty strings or range conditions.
- Improve documentation with more complex examples and use cases.
- Introduce integrations with logging frameworks to output detailed panics in production environments (if needed).

## Conclusion

With `must`, Go developers get a tool that prioritizes **assertive correctness**. Whether you're in the middle of a prototype or just testing, you don’t need to check everything manually. Let `must` take care of the basic sanity checks for you so you can focus on building great software.

For now, `must` might be small, but it’ll make sure your foundations are solid.

---

## Give stars

Feel free to contribute or improve the package! Your stars and pull requests are welcome.

## Thank You

If you find this package valuable, give it a star on GitHub! Thank you!!!
