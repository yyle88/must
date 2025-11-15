[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/must/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/must/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/must)](https://pkg.go.dev/github.com/yyle88/must)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/must/main.svg)](https://coveralls.io/github/yyle88/must?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23%2C%201.24%2C%201.25-lightgrey.svg)](https://github.com/yyle88/must)
[![GitHub Release](https://img.shields.io/github/release/yyle88/must.svg)](https://github.com/yyle88/must/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/must)](https://goreportcard.com/report/github.com/yyle88/must)

# must

简洁的断言工具包，提供 panic-on-failure 语义，减少 Go 代码样板。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🎯 **失败即崩溃验证**: 干净的断言机制，条件违背时自动触发 panic
⚡ **类型安全泛型**: 在所有断言类型上全面支持 Go 泛型
🔄 **栈帧调整**: 通过智能 skip 配置提供精确的 panic 位置
🌍 **结构化日志**: 与 zap 深度集成，提供详细的 panic 上下文
📋 **领域专用包**: 针对数值、字符串、切片和映射的专业工具

---

## 安装

```bash
go get github.com/yyle88/must
```

---

## 快速入门

### 示例 1: 断言非零值

```go
package main

import (
	"github.com/yyle88/must"
)

func main() {
	value := 42
	must.Nice(value) // 如果值为零，触发 panic

	println("值是有效的:", value)
}
```

---

### 示例 2: 验证没有错误

```go
package main

import (
	"errors"
	"github.com/yyle88/must"
)

func main() {
	err := someFunction()
	must.Done(err) // 如果 err 不为 nil，触发 panic

	println("没有遇到错误！")
}

func someFunction() error {
	return errors.New("意外的错误")
}
```

---

### 示例 3: 检查切片长度

```go
package main

import (
	"github.com/yyle88/must"
)

func main() {
	arr := []int{1, 2, 3}
	must.Length(arr, 3) // 如果长度不是 3，触发 panic

	println("切片长度正确")
}
```

---

## 核心断言

以下是 `must` 中的核心断言函数，概述如下：

| **函数**                       | **描述**                                           | **示例**                        | **备注**                 |
|------------------------------|--------------------------------------------------|-------------------------------|------------------------|
| **`True(v bool)`**           | 如果 `v` 为 `false`，触发 panic。                       | `must.True(isValid)`          | 验证 `v` 是否为 `true`。     |
| **`Done(err error)`**        | 如果 `err` 不为 `nil`，触发 panic。                      | `must.Done(err)`              | 确保没有错误发生。              |
| **`Must(err error)`**        | 如果 `err` 不为 `nil`，触发 panic。                      | `must.Must(err)`              | 类似于 `Done`。            |
| **`Nice(a V)`**              | 如果 `a` 为零，触发 panic。                              | `must.Nice(value)`            | 确保 `a` 非零。             |
| **`Zero(a V)`**              | 如果 `a` 不是零，触发 panic。                             | `must.Zero(value)`            | 确保 `a` 为零。             |
| **`None(a V)`**              | 如果 `a` 非零，触发 panic。                              | `must.None(value)`            | 确保 `a` 为零。             |
| **`Null(v any)`**            | 如果 `v` 不为 `nil`，触发 panic。                        | `must.Null(ptr)`              | 确保 `v` 为 `nil`。        |
| **`Full(v any)`**            | 如果 `v` 为 `nil`，触发 panic。                         | `must.Full(value)`            | 确保 `v` 非 `nil`。        |
| **`Equals(a, b V)`**         | 如果 `a` 和 `b` 不相等，触发 panic。                       | `must.Equals(a, b)`           | 检查 `a` 是否等于 `b`。       |
| **`Same(a, b V)`**           | 如果 `a` 和 `b` 不相等，触发 panic。                       | `must.Same(a, b)`             | `Equals` 的别名。          |
| **`SameNice(a, b V)`**       | 如果 `a` 和 `b` 不相等或为零，触发 panic。                  | `must.SameNice(a, b)`         | 确保相等且非零。             |
| **`Sane(a, b V)`**           | 如果 `a` 和 `b` 不相等或为零，触发 panic。                  | `must.Sane(a, b)`             | `SameNice` 的别名。        |
| **`Diff(a, b V)`**           | 如果 `a` 和 `b` 相等，触发 panic。                        | `must.Diff(a, b)`             | 确保值不同。               |
| **`Different(a, b V)`**      | 如果 `a` 和 `b` 相等，触发 panic。                        | `must.Different(a, b)`        | `Diff` 的别名。            |
| **`Is(a, b V)`**             | 如果 `a` 和 `b` 不相等，触发 panic。                       | `must.Is(a, b)`               | `Equals` 的别名。          |
| **`Ise(err, target error)`** | 如果 `err` 不与 `target` 匹配，触发 panic，使用 `errors.Is`。 | `must.Ise(err, targetErr)`    | 类似于 `errors.Is` 的错误匹配。 |
| **`Ok(a V)`**                | 如果 `a` 为零，触发 panic。                              | `must.Ok(value)`              | 确保 `a` 非零。             |
| **`OK(a V)`**                | `Ok` 的别名，检查值是否非零。                                | `must.OK(value)`              | 与 `Ok` 相同。             |
| **`TRUE(v bool)`**           | 如果 `v` 为 `false`，触发 panic。                       | `must.TRUE(isValid)`          | `True` 的别名。            |
| **`FALSE(v bool)`**          | 如果 `v` 为 `true`，触发 panic。                        | `must.FALSE(isError)`         | 确保 `v` 为 `false`。      |
| **`False(v bool)`**          | 如果 `v` 为 `true`，触发 panic。                        | `must.False(isError)`         | 与 `FALSE` 相同。          |
| **`Have(a []T)`**            | 如果 `a` 为空，触发 panic。                              | `must.Have(slice)`            | 确保 `a` 不为空。            |
| **`Length(a []T, n int)`**   | 如果 `a` 的长度不是 `n`，触发 panic。                       | `must.Length(slice, 3)`       | 确保 `a` 的长度是 `n`。       |
| **`Len(a []T, n int)`**      | `Length` 的别名，确保 `a` 的长度是 `n`。                    | `must.Len(slice, 3)`          | 验证 `a` 的长度。            |
| **`In(v T, a []T)`**         | 如果 `v` 不在 `a` 中，触发 panic。                        | `must.In(value, slice)`       | 确保 `v` 在 `a` 中。        |
| **`Contains(a []T, v T)`**   | 如果 `a` 不包含 `v`，触发 panic。                         | `must.Contains(slice, value)` | 确保 `a` 包含 `v`。         |

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yyle88/must.svg?variant=adaptive)](https://starchart.cc/yyle88/must)
