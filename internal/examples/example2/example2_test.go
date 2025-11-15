package example2

import (
	"testing"

	"github.com/yyle88/must"
)

// Example 2: Working with slices and collections
// Example 2: 切片和集合操作

func TestSliceHasItems(t *testing.T) {
	// Demo: Check slice has items
	// 演示：检查切片有元素
	items := []string{"apple", "banana", "orange"}
	must.Have(items)

	t.Logf("Found %d items", len(items))
}

func TestSliceWithExpectedSize(t *testing.T) {
	// Demo: Check slice has expected size
	// 演示：检查切片有预期大小
	numbers := []int{1, 2, 3, 4, 5}
	must.Length(numbers, 5)

	t.Log("Slice size is correct")
}

func TestSliceIncludesValue(t *testing.T) {
	// Demo: Check slice includes specific value
	// 演示：检查切片包含特定值
	fruits := []string{"apple", "banana", "orange"}
	must.Contains(fruits, "banana")

	t.Log("Found banana in the items")
}

func TestValueExistsInSlice(t *testing.T) {
	// Demo: Check value exists in slice
	// 演示：检查值存在于切片中
	numbers := []int{10, 20, 30, 40, 50}
	must.In(30, numbers)

	t.Log("Value 30 exists in the slice")
}

func TestMapOperations(t *testing.T) {
	// Demo: Working with maps
	// 演示：使用 map 操作
	config := map[string]int{
		"timeout": 30,
		"retries": 3,
		"port":    8080,
	}

	// Check map has content
	// 检查 map 有内容
	must.True(len(config) > 0)

	t.Logf("Config has %d entries", len(config))
}

func TestMultipleChecks(t *testing.T) {
	// Demo: Chain multiple checks
	// 演示：链式多个检查
	data := getData()

	must.True(data != nil)           // Data exists
	must.Have(data)                  // Data has items
	must.Length(data, 3)             // Data has 3 items
	must.Contains(data, "important") // Data contains item

	t.Log("Data checks passed")
}

func TestEmptySliceSafeHandling(t *testing.T) {
	// Demo: Safe handling when slice might be without items
	// 演示：安全处理可能无元素的切片
	items := getOptionalItems()

	if len(items) > 0 {
		must.Have(items) // Safe when we know it has items
		t.Logf("Processing %d items", len(items))
	} else {
		t.Log("No items to process")
	}
}

// Supporting functions used in tests
// 测试中使用的辅助函数

func getData() []string {
	return []string{"high", "important", "medium"}
}

func getOptionalItems() []string {
	// Simulate function that might return without items
	// 模拟可能返回无元素的函数
	return []string{}
}

// Advanced examples
// 高级示例

func TestFilteredData(t *testing.T) {
	// Demo: Work with processed/transformed data
	// 演示：处理过滤/转换后的数据
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Get even numbers
	// 获取偶数
	evenNumbers := getEvenNumbers(numbers)

	must.Have(evenNumbers)        // Has outcome
	must.Length(evenNumbers, 5)   // Correct count
	must.Contains(evenNumbers, 2) // Contains expected value
	must.In(4, evenNumbers)       // Expected value is present

	t.Logf("Got %d even numbers", len(evenNumbers))
}

func getEvenNumbers(numbers []int) []int {
	outcome := []int{}
	for _, n := range numbers {
		if n%2 == 0 {
			outcome = append(outcome, n)
		}
	}
	return outcome
}

func TestNestedData(t *testing.T) {
	// Demo: Working with nested data
	// 演示：处理嵌套数据
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	must.Have(matrix)         // Matrix has content
	must.Length(matrix, 3)    // Has 3 rows
	must.Have(matrix[0])      // Row has content
	must.Length(matrix[0], 3) // Row has 3 columns

	t.Log("Matrix validated")
}
