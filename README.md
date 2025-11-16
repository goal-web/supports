# Goal-Supports

[![Go Reference](https://pkg.go.dev/badge/github.com/goal-web/supports.svg)](https://pkg.go.dev/github.com/goal-web/supports)
[![Go Report Card](https://goreportcard.com/badge/github.com/goal-web/supports)](https://goreportcard.com/report/github.com/goal-web/supports)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](../goal/LICENSE)
![GitHub Stars](https://img.shields.io/github/stars/goal-web/supports?style=social)
![Release](https://img.shields.io/github/v/release/goal-web/supports?include_prereleases)
![Go Version](https://img.shields.io/badge/go-%3E=%201.25.0-00ADD8?logo=go)
![CI](https://img.shields.io/github/actions/workflow/status/goal-web/supports/ci.yml?branch=master&label=CI)
![Lint](https://img.shields.io/github/actions/workflow/status/goal-web/supports/lint.yml?branch=master&label=Lint)
![Commit Activity](https://img.shields.io/github/commit-activity/m/goal-web/supports)

[Docs](https://pkg.go.dev/github.com/goal-web/supports) · [Issues](https://github.com/goal-web/supports/issues) · [Releases](https://github.com/goal-web/supports/releases) · [中文文档](./README.cn.md)

Goal Supports provides common building blocks for Goal projects: type-safe field access (`BaseFields`), logging wrappers, command signature parsing, utilities, exceptions and signal providers.

## Highlights

- BaseFields: typed getters and optional defaults with provider pattern
- Logging: structured logging with debug switch
- Commands: DSL signature parsing, argument injection
- Utils: conversions, reflection, files, random, maps, fields flattening
- Exceptions/Signal: helpers and service providers

## Compatibility

- Go `>= 1.25.0`
- Module path: `github.com/goal-web/supports`

## Quick Start

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

## Commands

```go
sig := "demo {name} {count?} {--force} {--driver=aes}"
name, args := supports_commands.ParseSignature(sig)
_ = name
_ = args
```

## Star History

<a href="https://star-history.com/#goal-web/supports&Date"><img src="https://api.star-history.com/svg?repos=goal-web/supports&type=Date" alt="Star History Chart"/></a>

![Stargazers over time](https://starchart.cc/goal-web/supports.svg)

## CI Insights

- CI Status: ![CI](https://img.shields.io/github/actions/workflow/status/goal-web/supports/ci.yml?branch=master&label=CI)
- Lint Status: ![Lint](https://img.shields.io/github/actions/workflow/status/goal-web/supports/lint.yml?branch=master&label=Lint)
- Commit Activity: ![Commit Activity](https://img.shields.io/github/commit-activity/m/goal-web/supports)
- Last Commit: ![Last Commit](https://img.shields.io/github/last-commit/goal-web/supports)

## Contributing

- Keep API backward compatible, update docs/examples for changes
- Build and run tests before submitting
- Use Issues/Discussions for proposals

## Roadmap

- More tests for commands/exceptions/signal
- Example app and smoke tests
- English/Chinese docs maintained

## Changelog

- See Releases: <https://github.com/goal-web/supports/releases>

## Security

- Do not commit secrets; prefer env vars or encrypted configs

## FAQ

- Use typed getters to avoid implicit conversions
- Prefer local builds of tools aligned with target Go version
