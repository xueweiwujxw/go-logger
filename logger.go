package gologger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type LogFile struct {
	loggerF     *log.Logger
	logFile     *os.File
	debugMode   bool
	withFile    bool
	lock        *sync.Mutex
	initialized bool
}

var logger LogFile

func InitFileLoger(debugMode bool, withFile bool) {
	var err error
	logger.debugMode = debugMode
	logger.withFile = withFile
	logger.lock = &sync.Mutex{}
	if logger.withFile {
		file := "./" + time.Now().Format("2006-01-02") + ".log"
		logger.logFile, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("logfile init failed")
		}
		logger.loggerF = log.New(logger.logFile, "", log.LstdFlags)
	}
	logger.initialized = true
}

func CloseLogFile() {
	if !logger.initialized {
		return
	}
	logger.lock.Lock()
	defer logger.lock.Unlock()
	logger.logFile.Close()
}

func getCallerPosition() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "???"
	}
	funcName := runtime.FuncForPC(pc).Name()
	if idx := strings.LastIndex(funcName, "/"); idx >= 0 {
		funcName = funcName[idx+1:]
	}
	file = filepath.Base(file)
	return fmt.Sprintf("%s:%d %s", file, line, funcName)
}

func Info(v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[info] [%s] ", getCallerPosition()))
	log.Print(v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[info] [%s] ", getCallerPosition()))
		logger.loggerF.Print(v...)
	}
}

func Infof(format string, v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[info] [%s] ", getCallerPosition()))
	log.Printf(format, v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[info] [%s] ", getCallerPosition()))
		logger.loggerF.Printf(format, v...)
	}
}

func Infoln(v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[info] [%s] ", getCallerPosition()))
	log.Println(v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[info] [%s] ", getCallerPosition()))
		logger.loggerF.Println(v...)
	}
}

func Warn(v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[warn] [%s] ", getCallerPosition()))
	log.Print(v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[warn] [%s] ", getCallerPosition()))
		logger.loggerF.Print(v...)
	}
}

func Warnf(format string, v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[warn] [%s] ", getCallerPosition()))
	log.Printf(format, v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[warn] [%s] ", getCallerPosition()))
		logger.loggerF.Printf(format, v...)
	}
}

func Warnln(v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[warn] [%s] ", getCallerPosition()))
	log.Println(v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[warn] [%s] ", getCallerPosition()))
		logger.loggerF.Println(v...)
	}
}

func Error(v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[error] [%s] ", getCallerPosition()))
	log.Print(v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[error] [%s] ", getCallerPosition()))
		logger.loggerF.Print(v...)
	}
}

func Errorf(format string, v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[error] [%s] ", getCallerPosition()))
	log.Printf(format, v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[error] [%s] ", getCallerPosition()))
		logger.loggerF.Printf(format, v...)
	}
}

func Errorln(v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[error] [%s] ", getCallerPosition()))
	log.Println(v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[error] [%s] ", getCallerPosition()))
		logger.loggerF.Println(v...)
	}
}

func Fatal(v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[fatal] [%s] ", getCallerPosition()))
	log.Fatal(v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[fatal] [%s] ", getCallerPosition()))
		logger.loggerF.Fatal(v...)
	}
}

func Fatalf(format string, v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[fatal] [%s] ", getCallerPosition()))
	log.Fatalf(format, v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[fatal] [%s] ", getCallerPosition()))
		logger.loggerF.Fatalf(format, v...)
	}
}

func Fatalln(v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[fatal] [%s] ", getCallerPosition()))
	log.Fatalln(v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[fatal] [%s] ", getCallerPosition()))
		logger.loggerF.Fatalln(v...)
	}
}

func Panic(v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[panic] [%s] ", getCallerPosition()))
	log.Panic(v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[panic] [%s] ", getCallerPosition()))
		logger.loggerF.Panic(v...)
	}
}

func Panicf(format string, v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[panic] [%s] ", getCallerPosition()))
	log.Panicf(format, v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[panic] [%s] ", getCallerPosition()))
		logger.loggerF.Panicf(format, v...)
	}
}

func Panicln(v ...interface{}) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	log.SetPrefix(fmt.Sprintf("[panic] [%s] ", getCallerPosition()))
	log.Panicln(v...)
	if logger.withFile {
		logger.loggerF.SetPrefix(fmt.Sprintf("[panic] [%s] ", getCallerPosition()))
		logger.loggerF.Panicln(v...)
	}
}

func Debug(v ...interface{}) {
	if logger.debugMode {
		logger.lock.Lock()
		defer logger.lock.Unlock()
		log.SetPrefix(fmt.Sprintf("[debug] [%s] ", getCallerPosition()))
		log.Print(v...)
		if logger.withFile {
			logger.loggerF.SetPrefix(fmt.Sprintf("[debug] [%s] ", getCallerPosition()))
			logger.loggerF.Print(v...)
		}
	}
}

func Debugf(format string, v ...interface{}) {
	if logger.debugMode {
		logger.lock.Lock()
		defer logger.lock.Unlock()
		log.SetPrefix(fmt.Sprintf("[debug] [%s] ", getCallerPosition()))
		log.Printf(format, v...)
		if logger.withFile {
			logger.loggerF.SetPrefix(fmt.Sprintf("[debug] [%s] ", getCallerPosition()))
			logger.loggerF.Printf(format, v...)
		}
	}
}

func Debugln(v ...interface{}) {
	if logger.debugMode {
		logger.lock.Lock()
		defer logger.lock.Unlock()
		log.SetPrefix(fmt.Sprintf("[debug] [%s] ", getCallerPosition()))
		log.Println(v...)
		if logger.withFile {
			logger.loggerF.SetPrefix(fmt.Sprintf("[debug] [%s] ", getCallerPosition()))
			logger.loggerF.Println(v...)
		}
	}
}
