package example3

import (
	"testing"

	"github.com/yyle88/must"
)

// Example 3: Advanced patterns and scenarios
// Example 3: 高级模式和场景

// Account represents an account in the system
// Account 代表系统中的账号
type Account struct {
	ID       int
	Username string
	Mailbox  string
	Active   bool
}

// Config represents application configuration
// Config 代表应用配置
type Config struct {
	Host     string
	Port     int
	MaxConns int
	Options  map[string]string
}

func TestAccountValidation(t *testing.T) {
	// Demo: Validate account object
	// 演示：验证账号对象
	account := createAccount()

	must.Nice(account.ID)       // ID is set
	must.Nice(account.Username) // Username has content
	must.Nice(account.Mailbox)  // Mailbox has content
	must.True(account.Active)   // Account is active

	t.Logf("Account %s validated", account.Username)
}

func TestConfigSetup(t *testing.T) {
	// Demo: Validate configuration
	// 演示：验证配置
	cfg := loadConfig()

	must.Nice(cfg.Host)             // Host is configured
	must.Nice(cfg.Port)             // Port is set
	must.True(cfg.Port > 0)         // Port is positive
	must.True(cfg.MaxConns > 0)     // Max connections is positive
	must.True(len(cfg.Options) > 0) // Options map has content

	t.Logf("Config loaded: %s:%d", cfg.Host, cfg.Port)
}

func TestDataProcessing(t *testing.T) {
	// Demo: Process data with validation
	// 演示：处理数据并验证
	input := []int{5, 10, 15, 20, 25}

	// Get sum
	// 获取总和
	sum := getSum(input)
	must.Nice(sum)       // Sum is not zero
	must.True(sum > 0)   // Sum is positive
	must.Equals(sum, 75) // Sum is correct

	// Get average
	// 获取平均值
	avg := calculateAverage(input)
	must.Nice(avg)       // Average is not zero
	must.True(avg > 0)   // Average is positive
	must.Equals(avg, 15) // Average is correct

	t.Logf("Sum: %d, Average: %d", sum, avg)
}

func TestReferencesAndPointing(t *testing.T) {
	// Demo: Working with pointers
	// 演示：使用指针
	t.Run("Full pointer", func(t *testing.T) {
		value := 100
		ptr := &value

		must.Full(ptr)         // pointer is not absent
		must.Nice(*ptr)        // Value is not zero
		must.Equals(*ptr, 100) // Value is correct
	})

	t.Run("Null pointer", func(t *testing.T) {
		var ptr *int
		must.Null(ptr) // pointer is absent
	})

	t.Log("ptr checks passed")
}

func TestComparison(t *testing.T) {
	// Demo: Compare values
	// 演示：比较值
	a := 42
	b := 42
	c := 100

	must.Same(a, b)      // a and b are same
	must.Equals(a, b)    // a and b match
	must.Diff(a, c)      // a and c are distinct
	must.Different(b, c) // b and c are distinct

	t.Log("Comparison tests passed")
}

func TestChainedValidation(t *testing.T) {
	// Demo: Chain validations with complex logic
	// 演示：链式验证复杂逻辑
	outcome := processAndValidate()

	must.True(outcome != nil)    // Outcome exists
	must.True(len(outcome) > 0)  // Outcome has content
	must.Have(outcome)           // Outcome slice has items
	must.Contains(outcome, "ok") // Outcome contains success token

	t.Logf("Processing outcome: %v", outcome)
}

func TestResourceManagement(t *testing.T) {
	// Demo: Validate resource setup
	// 演示：验证资源配置
	resources := allocateResources(5)

	must.True(resources != nil) // Resources exist
	must.Length(resources, 5)   // Correct amount setup
	must.Have(resources)        // Has resources

	// Validate each resource
	// 验证每个资源
	for i, res := range resources {
		must.Nice(res) // Each resource is sound
		t.Logf("Resource %d: %s", i, res)
	}

	t.Log("Resources validated")
}

func TestComplexScenario(t *testing.T) {
	// Demo: Complex scenario
	// 演示：复杂场景
	account := createAccount()
	cfg := loadConfig()

	// Account setup
	// 账号设置
	must.Nice(account.ID)
	must.Nice(account.Username)
	must.True(account.Active)

	// Config validation
	// 配置验证
	must.Nice(cfg.Host)
	must.True(cfg.Port > 0)
	must.True(cfg.MaxConns >= 10)

	// Process account data with config
	// 使用配置处理账号数据
	outcome := processAccountWithConfig(account, cfg)
	must.Nice(outcome)
	must.True(outcome)

	t.Log("Complex scenario completed without issues")
}

// Supporting functions used in tests
// 测试中使用的辅助函数

func createAccount() Account {
	return Account{
		ID:       1001,
		Username: "alice",
		Mailbox:  "alice@example.com",
		Active:   true,
	}
}

func loadConfig() Config {
	return Config{
		Host:     "localhost",
		Port:     8080,
		MaxConns: 100,
		Options: map[string]string{
			"timeout": "30s",
			"retries": "3",
		},
	}
}

func getSum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func calculateAverage(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return getSum(numbers) / len(numbers)
}

func processAndValidate() []string {
	return []string{"ok", "processed", "validated"}
}

func allocateResources(count int) []string {
	resources := make([]string, count)
	for i := range resources {
		resources[i] = "resource-" + string(rune('A'+i))
	}
	return resources
}

func processAccountWithConfig(account Account, cfg Config) bool {
	// Simulate processing
	// 模拟处理
	return account.Active && cfg.Port > 0
}
