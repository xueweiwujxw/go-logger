<h2 align="center">
  Go logger
</h2>
<p align="center">
  Go logger is a logging library that supports writing log files. 
</p>
<p align="center">
  It is a wrapper around the standard Go log package and provides support for the following log levels: info, warn, error, debug, fatal, and panic.
</p>
<p align="center">
  <img src="https://github.com/xueweiwujxw/go-logger/actions/workflows/test.yml/badge.svg" alt="Test">
  <img src="https://github.com/xueweiwujxw/go-logger/actions/workflows/release.yml/badge.svg" alt="Release">
  <br/>
  <a href="https://codecov.io/gh/xueweiwujxw/go-logger">
    <img src="https://codecov.io/gh/xueweiwujxw/go-logger/branch/master/graph/badge.svg?token=5WBX2UZQ79" alt="Codecov">
  </a>
  <a href="https://pkg.go.dev/github.com/xueweiwujxw/go-logger">
    <img src="https://pkg.go.dev/badge/github.com/xueweiwujxw/go-logger.svg" alt="Go Reference">
  </a>
  <a href="https://github.com/xueweiwujxw/go-logger/releases">
    <img src="https://img.shields.io/github/v/tag/xueweiwujxw/go-logger" alt="GitHub tag">
  </a>
  <img src="https://img.shields.io/github/repo-size/xueweiwujxw/multi-provider" alt="repo size">
  <a href="https://app.codacy.com/gh/xueweiwujxw/go-logger/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade">
    <img src="https://app.codacy.com/project/badge/Grade/bf2e5a8ab0c544f5a65e122220df0f06" alt="Codacy Badge">
  </a>
</p>
<hr/>

### Installation

```shell
go get github.com/xueweiwujxw/go-logger
```

### Usage

```go
package main

import (
	gologger "github.com/xueweiwujxw/go-logger"
)

func main() {
	gologger.InitFileLoger(true, true, "")

	gologger.Info("Info test\n")
	gologger.Infof("%s\n", "Infof test")
	gologger.Infoln("Infoln test")

	gologger.Warn("Warn test\n")
	gologger.Warnf("%s\n", "Warnf test")
	gologger.Warnln("Warnln test")

	gologger.Error("Error test\n")
	gologger.Errorf("%s\n", "Errorf test")
	gologger.Errorln("Errorln test")

	gologger.Debug("Debug test\n")
	gologger.Debugf("%s\n", "Debugf test")
	gologger.Debugln("Debugln test")

	gologger.CloseLogFile()
}
```

```shell
2023/05/19 12:04:59 [info] [usage.go:10 main.main] Info test
2023/05/19 12:04:59 [info] [usage.go:11 main.main] Infof test
2023/05/19 12:04:59 [info] [usage.go:12 main.main] Infoln test
2023/05/19 12:04:59 [warn] [usage.go:14 main.main] Warn test
2023/05/19 12:04:59 [warn] [usage.go:15 main.main] Warnf test
2023/05/19 12:04:59 [warn] [usage.go:16 main.main] Warnln test
2023/05/19 12:04:59 [error] [usage.go:18 main.main] Error test
2023/05/19 12:04:59 [error] [usage.go:19 main.main] Errorf test
2023/05/19 12:04:59 [error] [usage.go:20 main.main] Errorln test
2023/05/19 12:04:59 [debug] [usage.go:22 main.main] Debug test
2023/05/19 12:04:59 [debug] [usage.go:23 main.main] Debugf test
2023/05/19 12:04:59 [debug] [usage.go:24 main.main] Debugln test
```

### Contributors

<a href="https://github.com/xueweiwujxw/go-logger/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=xueweiwujxw/go-logger" />
</a>
