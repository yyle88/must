package main

import (
	"fmt"

	"github.com/yyle88/must"
)

// Demo2x demonstrates rese package functions
// Demo2x 演示 rese 包函数
func main() {
	fmt.Println("=== Demo 2: Rese Package ===")

	// V1 - single value validation
	config := must.V1(readConfig())
	fmt.Printf("✓ Config: %s\n", config)

	// V2 - dual value validation
	width, height := must.V2(getDimensions())
	fmt.Printf("✓ Dimensions: %dx%d\n", width, height)

	// P1 - pointer validation
	admin := must.P1(findAdmin())
	fmt.Printf("✓ Admin: %s\n", admin.Name)

	// C1 - non-zero validation
	num := must.C1(getNum())
	fmt.Printf("✓ Num: %d\n", num)

	// Combined validations
	data := getData()
	must.Full(data)
	must.Nice(data.Score)
	must.Same(data.Status, "active")
	fmt.Printf("✓ Data: score=%d, status=%s\n", data.Score, data.Status)

	fmt.Println("\n=== All checks passed! ===")
}

type Admin struct{ Name string }
type Data struct {
	Score  int
	Status string
}

func readConfig() (string, error)      { return "v1.0", nil }
func getDimensions() (int, int, error) { return 1920, 1080, nil }
func findAdmin() (*Admin, error)       { return &Admin{"Alice"}, nil }
func getNum() (int, error)             { return 123, nil }
func getData() *Data                   { return &Data{95, "active"} }
