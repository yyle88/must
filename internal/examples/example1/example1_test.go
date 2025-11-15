package example1

import (
	"fmt"
	"testing"

	"github.com/yyle88/must"
)

// Example 1: Basic value checks and operation validation
// Example 1: 基础值检查和操作验证

func TestBasicOperationCheck(t *testing.T) {
	// Demo: Check operation completes without issues
	// 演示：检查操作完成无问题
	issue := performOperation()
	must.Done(issue)

	t.Log("Operation completed without issues")
}

func TestValueAssertion(t *testing.T) {
	// Demo: Check non-zero value
	// 演示：检查非零值
	outcome := calculateOutcome()
	must.Nice(outcome)

	t.Logf("Got valid outcome: %d", outcome)
}

func TestConditionCheck(t *testing.T) {
	// Demo: Check condition passes
	// 演示：检查条件通过
	passes := checkCondition()
	must.True(passes)

	t.Log("Condition check passed")
}

func TestMatchCheck(t *testing.T) {
	// Demo: Check two values match
	// 演示：检查两个值匹配
	expected := "success"
	outcome := getOutcome()
	must.Equals(expected, outcome)

	t.Log("Status check passed")
}

// Supporting functions used in tests
// 测试中使用的辅助函数

func performOperation() (issue error) {
	// Simulate operation that completes without issues
	// 模拟无问题完成的操作
	return nil
}

func calculateOutcome() int {
	// Simulate computation
	// 模拟计算
	return 42
}

func checkCondition() bool {
	// Simulate condition check
	// 模拟条件检查
	return true
}

func getOutcome() string {
	// Simulate status check
	// 模拟状态检查
	return "success"
}

// Negative examples (commented out to avoid panic in tests)
// 反面示例（注释掉以避免测试中的 panic）

/*
func TestOperationWithIssue(t *testing.T) {
	// This would panic because operation has issue
	// 这会触发 panic 因为操作有问题
	issue := fmt.Errorf("something went wrong")
	must.Done(issue) // PANIC!
}

func TestZeroOutcomeFails(t *testing.T) {
	// This would panic because outcome is zero
	// 这会触发 panic 因为结果为零
	outcome := 0
	must.Nice(outcome) // PANIC!
}
*/

func TestOperationWithIssueHandled(t *testing.T) {
	// Demo: Show safe handling pattern
	// 演示：展示安全处理模式
	issue := performRiskyTask()
	if issue != nil {
		t.Logf("Task has issue as expected: %v", issue)
	} else {
		must.Done(issue) // Safe when we know it's nil
	}
}

func performRiskyTask() (issue error) {
	// Simulate task that might have issues
	// 模拟可能有问题的任务
	return fmt.Errorf("simulated issue")
}
