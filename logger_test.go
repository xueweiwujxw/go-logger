package gologger_test

import (
	"bufio"
	"os"
	"strings"
	"sync"
	"testing"

	gologger "github.com/xueweiwujxw/go-logger"
)

func TestGologger(t *testing.T) {
	expected := []string{
		"[info]", " Info test",
		"[info]", " Infof test",
		"[info]", " Infoln test",
		"[warn]", " Warn test",
		"[warn]", " Warnf test",
		"[warn]", " Warnln test",
		"[error]", " Error test",
		"[error]", " Errorf test",
		"[error]", " Errorln test",
		"[debug]", " Debug test",
		"[debug]", " Debugf test",
		"[debug]", " Debugln test",
	}

	oldStderr := os.Stderr
	stderrR, stderrW, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stderr = stderrW

	var b []string
	var stderrWaitGroup sync.WaitGroup
	stderrWaitGroup.Add(1)
	go func() {
		defer stderrWaitGroup.Done()
		scanner := bufio.NewScanner(stderrR)
		for scanner.Scan() {
			b = append(b, scanner.Text())
		}
	}()

	// Test code begin
	gologger.InitFileLoger(true, false)

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
	// Test code end

	stderrW.Close()
	os.Stderr = oldStderr

	stderrWaitGroup.Wait()

	if len(b) < len(expected)/2 {
		for _, v := range b {
			t.Errorf("%q", v)
		}
		t.Fatalf("no enough output, got %d, expected %d", len(b), len(expected)/2)
	}

	for i, v := range b {
		if !strings.Contains(v, expected[i*2]) {
			t.Errorf("log error, got %q, does not contain expected %q", v, expected[i*2])
		}
		if !strings.Contains(v, expected[i*2+1]) {
			t.Errorf("log error, got %q, does not contain expected %q", v, expected[i*2+1])
		}
	}
}

func TestGologgerWithFile(t *testing.T) {
	expected := []string{
		"[info]", " Info test",
		"[info]", " Infof test",
		"[info]", " Infoln test",
		"[warn]", " Warn test",
		"[warn]", " Warnf test",
		"[warn]", " Warnln test",
		"[error]", " Error test",
		"[error]", " Errorf test",
		"[error]", " Errorln test",
		"[debug]", " Debug test",
		"[debug]", " Debugf test",
		"[debug]", " Debugln test",
	}

	oldStderr := os.Stderr
	stderrR, stderrW, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stderr = stderrW

	var b []string
	var stderrWaitGroup sync.WaitGroup
	stderrWaitGroup.Add(1)
	go func() {
		defer stderrWaitGroup.Done()
		scanner := bufio.NewScanner(stderrR)
		for scanner.Scan() {
			// do nothing
		}
	}()

	// Test code begin
	gologger.InitFileLoger(true, true)

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
	// Test code end

	stderrW.Close()
	os.Stderr = oldStderr

	stderrWaitGroup.Wait()

	file, err := os.OpenFile(gologger.GetFileName(), os.O_RDONLY, 0666)
	if err != nil {
		t.Fatalf("log file open failed, %q", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b = append(b, scanner.Text())
	}

	file.Close()
	os.Remove(gologger.GetFileName())

	if len(b) < len(expected)/2 {
		for _, v := range b {
			t.Errorf("%q", v)
		}
		t.Fatalf("no enough output, got %d, expected %d", len(b), len(expected)/2)
	}

	for i, v := range b {
		if !strings.Contains(v, expected[i*2]) {
			t.Errorf("log error, got %q, does not contain expected %q", v, expected[i*2])
		}
		if !strings.Contains(v, expected[i*2+1]) {
			t.Errorf("log error, got %q, does not contain expected %q", v, expected[i*2+1])
		}
	}

}
