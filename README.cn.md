# Goal-Supports

[![Go Reference](https://pkg.go.dev/badge/github.com/goal-web/supports.svg)](https://pkg.go.dev/github.com/goal-web/supports)
[![Go Report Card](https://goreportcard.com/badge/github.com/goal-web/supports)](https://goreportcard.com/report/github.com/goal-web/supports)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](../goal/LICENSE)
![GitHub Stars](https://img.shields.io/github/stars/goal-web/supports?style=social)
![Release](https://img.shields.io/github/v/release/goal-web/supports?include_prereleases)
![Go Version](https://img.shields.io/badge/go-%3E=%201.25.0-00ADD8?logo=go)
![CI](https://img.shields.io/github/actions/workflow/status/goal-web/supports/ci.yml?branch=master&label=CI)
![Lint](https://img.shields.io/github/actions/workflow/status/goal-web/supports/lint.yml?branch=master&label=Lint)

[Docs](https://pkg.go.dev/github.com/goal-web/supports) · [Issues](https://github.com/goal-web/supports/issues) · [Releases](https://github.com/goal-web/supports/releases) · [English](./README.md)

Supports 组件为 Goal 项目提供通用能力：类型安全的字段读取（`BaseFields`）、日志封装、命令签名解析、工具集合、异常与信号服务提供者。

## 亮点

- BaseFields：类型化读取与可选默认值，Provider 模式
- 日志：结构化日志与 Debug 开关
- 命令：DSL 签名解析与参数注入
- 工具：转换、反射、文件、随机、字段展平
- 异常/信号：辅助与服务提供者

## 兼容性

- Go `>= 1.25.0`
- 模块路径：`github.com/goal-web/supports`

## 快速开始

```go
package main

import (
    "github.com/goal-web/supports"
    "github.com/goal-web/supports/logs"
    "github.com/goal-web/supports/utils"
)

type provider struct{ supports.BaseFields }

func main() {
    p := &provider{}
    p.OptionalGetter = func(key string, defaultValue any) any { return defaultValue }

    v := p.StringOptional("app.name", "goal")
    logs.Default().WithField("name", v).Info("ready")

    f := utils.ToFields(map[string]any{"a": map[string]any{"b": 1}})
    utils.Flatten(make(map[string]any), f, ".")
}
```

## 命令签名

```go
sig := "demo {name} {count?} {--force} {--driver=aes}"
name, args := supports_commands.ParseSignature(sig)
_ = name
_ = args
```

## Star History

<a href="https://star-history.com/#goal-web/supports&Date"><img src="https://api.star-history.com/svg?repos=goal-web/supports&type=Date" alt="Star History Chart"/></a>

![Stargazers over time](https://starchart.cc/goal-web/supports.svg)

## 贡献

- 保持 API 向后兼容，更新文档与示例
- 提交前构建并运行测试
- 建议与讨论请使用 Issues/Discussions

## 路线图

- 补充 commands/exceptions/signal 测试
- 示例与冒烟测试
- 维护双语文档

## 变更日志

- 见 Releases：<https://github.com/goal-web/supports/releases>

## 安全

- 不要提交敏感信息；优先环境变量或加密配置

## FAQ

- 使用类型化读取，避免隐式转换
- 确保工具链与目标 Go 版本一致

