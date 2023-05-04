![Test](https://github.com/xueweiwujxw/go-logger/actions/workflows/test.yml/badge.svg)
![release](https://github.com/xueweiwujxw/go-logger/actions/workflows/release.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/xueweiwujxw/go-logger.svg)](https://pkg.go.dev/github.com/xueweiwujxw/go-logger)

## Go logger

Go logger is a logging library that supports writing log files. It is a wrapper around the standard Go log package and provides support for the following log levels: info, warn, error, debug, fatal, and panic.

** *

### Usage

```go
package main

import (
	gologger "github.com/xueweiwujxw/go-logger"
)

func main() {
	gologger.InitFileLoger(false, true)

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
2023/05/04 16:29:13 [info] [usage.go:10 main.main] Info test
2023/05/04 16:29:13 [info] [usage.go:11 main.main] Infof test
2023/05/04 16:29:13 [info] [usage.go:12 main.main] Infoln test
2023/05/04 16:29:13 [warn] [usage.go:14 main.main] Warn test
2023/05/04 16:29:13 [warn] [usage.go:15 main.main] Warnf test
2023/05/04 16:29:13 [warn] [usage.go:16 main.main] Warnln test
2023/05/04 16:29:13 [error] [usage.go:18 main.main] Error test
2023/05/04 16:29:13 [error] [usage.go:19 main.main] Errorf test
2023/05/04 16:29:13 [error] [usage.go:20 main.main] Errorln test
```

### Supported Functions

| func          |
| ------------- |
| InitFileLoger |
| GetFileName   |
| CloseLogFile  |
| Info          |
| Infof         |
| Infoln        |
| Warn          |
| Warnf         |
| Warnln        |
| Error         |
| Errorf        |
| Errorln       |
| Debug         |
| Debugf        |
| Debugln       |
| Fatal         |
| Fatalf        |
| Fatalln       |
| Panic         |
| Panicf        |
| Panicln       |


### Contributors

<a href="https://github.com/xueweiwujxw/go-logger/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=xueweiwujxw/go-logger" />
</a>
