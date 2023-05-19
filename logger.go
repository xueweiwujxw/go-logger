// @author wlanxww (xueweiwujxw@outlook.com)
// @version 0.1.4

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
	console     *log.Logger
	file        *log.Logger
	filename    string
	logFile     *os.File
	debugMode   bool
	withFile    bool
	lock        *sync.Mutex
	initialized bool
}

var logger LogFile
var OsExit = os.Exit

type messageLevel = int

const (
	infoLevel messageLevel = iota
	warnLevel
	errorLevel
	fatalLevel
	panicLevel
	debugLevel
)

var prefix = []string{"[info]", "[warn]", "[error]", "[fatal]", "[panic]", "[debug]"}

func (l *LogFile) output(level messageLevel, message string) {
	if !l.initialized {
		return
	}
	l.lock.Lock()
	defer l.lock.Unlock()

	if !l.debugMode && level == debugLevel {
		return
	}

	prefix := fmt.Sprintf("%s [%s] ", prefix[level], getCallerPosition())

	formatted := prefix + message

	l.console.Output(3, formatted)
	if l.withFile {
		l.file.Output(3, formatted)
	}
}

// SwitchExit is used to switch the program exit function.
// The `exit` parameter is a function that accepts an integer argument and replaces the standard os.Exit function.
//
// Note: This function is intended to be used only in a testing environment and should not be used in non-testing environments.
// In a testing environment, SwitchExit can be used to replace the program exit function with a custom function to capture the exit behavior during tests.
// Make sure that the custom exit function follows the behavior conventions of os.Exit, where the status code should be within the range [0, 125].
// It is not recommended to use this function to replace the os.Exit function in non-testing environments to avoid unnecessary side effects and erroneous behavior.
func SwitchExit(exit func(int)) {
	OsExit = exit
}

// Init File logger
//
//	@param debugMde bool enable debug log output
//	@param withFile bool enable log file generation
//	@param logFileName string custom log file, disabled when value is "" - empty string
//
// The default filename is the current date, and subsequent logs will be appended to the file.
// Custom filename functionality may be added in the future.
func InitFileLoger(debugMode bool, withFile bool, logFileName string) {
	var err error
	logger.debugMode = debugMode
	logger.withFile = withFile
	logger.lock = &sync.Mutex{}
	logger.console = log.New(os.Stderr, "", log.LstdFlags)
	if logger.withFile {
		var file string
		if logFileName == "" {
			file = "./" + time.Now().Format("2006-01-02") + ".log"
		} else {
			file = logFileName
		}
		logger.filename = file
		logger.logFile, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("logfile init failed")
		}
		logger.file = log.New(logger.logFile, "", log.LstdFlags)
	} else {
		logger.filename = ""
	}
	logger.initialized = true
}

// Get log file name if enabled
//
//	@returns string
func GetFileName() string {
	return logger.filename
}

// Close log file
func CloseLogFile() {
	if !logger.initialized {
		return
	}
	logger.lock.Lock()
	defer logger.lock.Unlock()
	logger.logFile.Close()
}

func getCallerPosition() string {
	pc, file, line, ok := runtime.Caller(3)
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
	logger.output(infoLevel, fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	logger.output(infoLevel, fmt.Sprintf(format, v...))
}

func Infoln(v ...interface{}) {
	logger.output(infoLevel, fmt.Sprintln(v...))
}

func Warn(v ...interface{}) {
	logger.output(warnLevel, fmt.Sprint(v...))
}

func Warnf(format string, v ...interface{}) {
	logger.output(warnLevel, fmt.Sprintf(format, v...))
}

func Warnln(v ...interface{}) {
	logger.output(warnLevel, fmt.Sprintln(v...))
}

func Error(v ...interface{}) {
	logger.output(errorLevel, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	logger.output(errorLevel, fmt.Sprintf(format, v...))
}

func Errorln(v ...interface{}) {
	logger.output(errorLevel, fmt.Sprintln(v...))
}

func Fatal(v ...interface{}) {
	logger.output(fatalLevel, fmt.Sprint(v...))
	CloseLogFile()
	OsExit(1)
}

func Fatalf(format string, v ...interface{}) {
	logger.output(fatalLevel, fmt.Sprintf(format, v...))
	CloseLogFile()
	OsExit(1)
}

func Fatalln(v ...interface{}) {
	logger.output(fatalLevel, fmt.Sprintln(v...))
	CloseLogFile()
	OsExit(1)
}

func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	logger.output(panicLevel, s)
	CloseLogFile()
	panic(s)
}

func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	logger.output(panicLevel, s)
	CloseLogFile()
	panic(s)
}

func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	logger.output(panicLevel, s)
	CloseLogFile()
	panic(s)
}

func Debug(v ...interface{}) {
	logger.output(debugLevel, fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	logger.output(debugLevel, fmt.Sprintf(format, v...))
}

func Debugln(v ...interface{}) {
	logger.output(debugLevel, fmt.Sprintln(v...))
}
