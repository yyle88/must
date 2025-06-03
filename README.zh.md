# `must` – Go 简单断言库

在 Go 中，错误处理可能显得繁琐。`must` 提供断言，在失败时触发 panic，在接收的返回值不符合预期时报错，这样能避免写很长的错误检查代码。

---

## 英文文档

[English README](README.md)

## 特性

- **简单的错误处理**：直接断言条件，避免冗长的检查代码。
- **快速反馈**：通过清晰的 panic 信息，及早捕捉 bug。
- **轻量且高效**：最小化的代码开销，提升速度，易于使用。
- **多样化断言**：支持检查非零值、切片长度等。

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

	println("数组长度正确")
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

## 许可证类型

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

---

## 贡献新代码

非常欢迎贡献代码！贡献流程：

1. 在 GitHub 上 Fork 仓库 （通过网页界面操作）。
2. 克隆Forked项目 (`git clone https://github.com/yourname/repo-name.git`)。
3. 在克隆的项目里 (`cd repo-name`)
4. 创建功能分支（`git checkout -b feature/xxx`）。
5. 添加代码 (`git add .`)。
6. 提交更改（`git commit -m "添加功能 xxx"`）。
7. 推送分支（`git push origin feature/xxx`）。
8. 发起 Pull Request （通过网页界面操作）。

请确保测试通过并更新相关文档。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
