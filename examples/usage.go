package main

import (
	gologger "github.com/xueweiwujxw/go-logger"
)

func main() {
	gologger.InitFileLoger(false, true)
	gologger.Infoln("Info test")
	gologger.CloseLogFile()
}
