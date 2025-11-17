package main

import (
	"fmt"

	"github.com/yyle88/must"
	"github.com/yyle88/must/mustmap"
	"github.com/yyle88/must/mustnum"
	"github.com/yyle88/must/mustslice"
	"github.com/yyle88/must/muststrings"
)

// Demo3x demonstrates specialized domain packages
// Demo3x 演示专用领域包
func main() {
	fmt.Println("=== Demo 3: Advanced Packages ===")

	// Numeric validations
	score := getScore()
	mustnum.Positive(score)
	mustnum.Gt(score, 60)
	fmt.Printf("✓ Score: %d\n", score)

	// Slice validations
	tags := getTags()
	mustslice.Have(tags)
	mustslice.Contains(tags, "go")
	fmt.Printf("✓ Tags: %v\n", tags)

	// Map validations
	config := getConfig()
	mustmap.Have(config)
	timeout := mustmap.Get(config, "timeout")
	fmt.Printf("✓ Timeout: %d\n", timeout)

	// String validations
	filename := getFilename()
	muststrings.HasSuffix(filename, ".pdf")
	muststrings.Contains(filename, "report")
	fmt.Printf("✓ Filename: %s\n", filename)

	// Complex scenario
	data := getAnalytics()
	must.Full(data)
	mustmap.Have(data.Metrics)
	fmt.Printf("✓ Analytics: %d metrics\n", len(data.Metrics))

	fmt.Println("\n=== All checks passed! ===")
}

type Analytics struct {
	Metrics map[string]float64
}

func getScore() int             { return 85 }
func getTags() []string         { return []string{"go", "test"} }
func getConfig() map[string]int { return map[string]int{"timeout": 30} }
func getFilename() string       { return "report.pdf" }
func getAnalytics() *Analytics {
	return &Analytics{Metrics: map[string]float64{"score": 87.5}}
}
