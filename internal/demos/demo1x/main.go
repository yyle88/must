package main

import (
	"fmt"

	"github.com/yyle88/must"
)

// Demo1x demonstrates basic assertion functions
// Demo1x 演示基础断言函数
func main() {
	fmt.Println("=== Demo 1: Basic Assertions ===")

	// Boolean assertion
	must.True(checkCondition())
	fmt.Println("✓ Boolean check passed")

	// Error check
	must.Done(performOperation())
	fmt.Println("✓ No error")

	// Non-zero value
	count := getCount()
	must.Nice(count)
	fmt.Printf("✓ Valid count: %d\n", count)

	// Equality check
	must.Equals("success", getStatus())
	fmt.Println("✓ Values match")

	// Slice operations
	items := getItems()
	must.Have(items)
	must.Length(items, 3)
	must.In("banana", items)
	fmt.Printf("✓ Slice validated: %v\n", items)

	// Pointer check
	account := getAccount()
	must.Full(account)
	fmt.Printf("✓ Pointer valid: %s\n", account.Name)

	fmt.Println("\n=== All checks passed! ===")
}

type Account struct{ Name string }

func checkCondition() bool    { return true }
func performOperation() error { return nil }
func getStatus() string       { return "success" }
func getCount() int           { return 42 }
func getItems() []string      { return []string{"apple", "banana", "orange"} }
func getAccount() *Account    { return &Account{Name: "test"} }
