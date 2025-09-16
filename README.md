# cppstring

## 这个仓库是用来干什么的 (Repository Purpose)

`cppstring` 是一个用 Go 语言编写的字符串操作库，提供类似 C++ 风格的字符串处理功能。尽管名称暗示与 C++ 相关，但这实际上是一个纯 Go 语言包，旨在提供高效的字符串操作工具。

### 主要功能 (Main Features)

- **字符串反转 (String Reversal)**: 提供高效的字符串反转功能，正确处理 Unicode 字符
- **模块化设计 (Modular Design)**: 采用清晰的包结构，便于扩展和维护
- **Go 模块支持 (Go Module Support)**: 完全支持 Go 模块系统，便于依赖管理

### 仓库结构 (Repository Structure)

```
cppstring/
├── cppstring.go        # 主包文件，导出公共 API
├── cppstring_test.go   # 测试文件
├── go.mod              # Go 模块定义
├── helpers/            # 辅助函数包
│   └── index.go        # 具体实现函数
├── examples/           # 使用示例
│   ├── example.go      # 示例程序
│   └── go.mod          # 示例模块配置
└── .gitignore          # Git 忽略文件配置
```

### 安装和使用 (Installation and Usage)

#### 安装 (Installation)

```bash
go get github.com/West-Pavilion/cppstring
```

#### 使用示例 (Usage Example)

```go
package main

import (
    "fmt"
    "github.com/West-Pavilion/cppstring"
)

func main() {
    // 字符串反转示例
    original := "Hello, 世界!"
    reversed := cppstring.Reverse(original)
    fmt.Printf("原字符串: %s\n", original)
    fmt.Printf("反转后: %s\n", reversed)
    // 输出: !界世 ,olleH
}
```

### API 文档 (API Documentation)

#### `Reverse(s string) string`

反转给定的字符串，正确处理 Unicode 字符（包括中文、表情符号等）。

**参数:**
- `s string`: 需要反转的字符串

**返回值:**
- `string`: 反转后的字符串

**示例:**
```go
result := cppstring.Reverse("ABC")  // 返回 "CBA"
result := cppstring.Reverse("你好")  // 返回 "好你"
```

### 开发信息 (Development Information)

- **语言版本 (Go Version)**: Go 1.20+
- **模块路径 (Module Path)**: `github.com/West-Pavilion/cppstring`
- **许可证 (License)**: 请查看仓库许可证文件

### 构建和测试 (Build and Test)

```bash
# 构建项目
go build

# 整理依赖
go mod tidy

# 运行测试
go test -v

# 运行性能测试
go test -bench=.

# 运行示例程序
cd examples && go run example.go
```

#### 测试结果 (Test Results)

所有测试均通过，包括：
- 空字符串处理
- 单字符字符串
- 基本英文字符串
- 含空格的字符串
- 中文字符
- 混合字符（英文+中文+数字）
- 表情符号
- 回文字符串

性能基准测试显示平均处理时间约为 635.8 ns/op。

### 贡献 (Contributing)

欢迎贡献代码！请确保：
1. 代码符合 Go 语言规范
2. 添加适当的测试用例
3. 更新相关文档

### 未来计划 (Future Plans)

- 添加更多字符串操作函数
- 增加性能优化
- 添加完整的测试套件
- 支持更多 C++ 风格的字符串操作

---

## English

`cppstring` is a Go library that provides C++-style string manipulation functions. Despite the name suggesting C++, this is actually a pure Go package designed to offer efficient string operation tools.

### Key Features

- **String Reversal**: Efficient string reversal with proper Unicode character handling
- **Modular Design**: Clean package structure for easy extension and maintenance  
- **Go Module Support**: Full Go module system support for easy dependency management

### Installation

```bash
go get github.com/West-Pavilion/cppstring
```

### Usage

```go
package main

import (
    "fmt"
    "github.com/West-Pavilion/cppstring"
)

func main() {
    original := "Hello, World!"
    reversed := cppstring.Reverse(original)
    fmt.Printf("Original: %s\n", original)
    fmt.Printf("Reversed: %s\n", reversed)
    // Output: !dlroW ,olleH
}
```