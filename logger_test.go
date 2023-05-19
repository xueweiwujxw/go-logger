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
	gologger.InitFileLoger(true, false, "")

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

func TestGologgerWithoutInitialize(t *testing.T) {
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

	if len(b) > 0 {
		for _, v := range b {
			t.Errorf("%q", v)
		}
		t.Fatalf("error output, got %d, expected empty", len(b))
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

func TestGologgerWithCustomFile(t *testing.T) {
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
	gologger.InitFileLoger(true, true, "./custom.log")

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

func TestFatal(t *testing.T) {
	testlog := "testFatal.log"
	expected := "Fatal test"

	oldOsExit := gologger.OsExit
	var gotExitCode int
	TestExit := func(code int) {
		gotExitCode = code
	}
	gologger.SwitchExit(TestExit)

	oldStderr := os.Stderr
	_, stderrW, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stderr = stderrW

	// test being
	gologger.InitFileLoger(false, true, testlog)
	gologger.Fatal("Fatal test\n")
	// test end

	gologger.SwitchExit(oldOsExit)

	stderrW.Close()
	os.Stderr = oldStderr

	if exp := 1; gotExitCode != exp {
		t.Errorf("expected exit code: %d, got: %d", exp, gotExitCode)
	}

	file, err := os.OpenFile(testlog, os.O_RDONLY, 0666)
	if err != nil {
		t.Fatalf("log file open failed, %q", err)
	}

	var b []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b = append(b, scanner.Text())
	}
	file.Close()
	os.Remove(testlog)

	if len(b) != 1 {
		t.Fatalf("invalid not output len, got %d, expected 1", len(b))
	}
	if !strings.Contains(b[0], expected) {
		t.Errorf("log error, got %q, does not cotain expected %q", b[0], expected)
	}
}

func TestFatalf(t *testing.T) {
	testlog := "testFatalf.log"
	expected := "Fatalf test"

	oldOsExit := gologger.OsExit
	var gotExitCode int
	TestExit := func(code int) {
		gotExitCode = code
	}
	gologger.SwitchExit(TestExit)

	oldStderr := os.Stderr
	_, stderrW, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stderr = stderrW

	// test being
	gologger.InitFileLoger(false, true, testlog)
	gologger.Fatalf("%s\n", "Fatalf test")
	// test end

	gologger.SwitchExit(oldOsExit)

	stderrW.Close()
	os.Stderr = oldStderr

	if exp := 1; gotExitCode != exp {
		t.Errorf("expected exit code: %d, got: %d", exp, gotExitCode)
	}

	file, err := os.OpenFile(testlog, os.O_RDONLY, 0666)
	if err != nil {
		t.Fatalf("log file open failed, %q", err)
	}

	var b []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b = append(b, scanner.Text())
	}
	file.Close()
	os.Remove(testlog)

	if len(b) != 1 {
		t.Fatalf("invalid not output len, got %d, expected 1", len(b))
	}
	if !strings.Contains(b[0], expected) {
		t.Errorf("log error, got %q, does not cotain expected %q", b[0], expected)
	}
}

func TestFatalln(t *testing.T) {
	testlog := "testFatalln.log"
	expected := "Fatalln test"

	oldOsExit := gologger.OsExit
	var gotExitCode int
	TestExit := func(code int) {
		gotExitCode = code
	}
	gologger.SwitchExit(TestExit)

	oldStderr := os.Stderr
	_, stderrW, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stderr = stderrW

	// test being
	gologger.InitFileLoger(false, true, testlog)
	gologger.Fatalln("Fatalln test")
	// test end

	gologger.SwitchExit(oldOsExit)

	stderrW.Close()
	os.Stderr = oldStderr

	if exp := 1; gotExitCode != exp {
		t.Errorf("expected exit code: %d, got: %d", exp, gotExitCode)
	}

	file, err := os.OpenFile(testlog, os.O_RDONLY, 0666)
	if err != nil {
		t.Fatalf("log file open failed, %q", err)
	}

	var b []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b = append(b, scanner.Text())
	}
	file.Close()
	os.Remove(testlog)

	if len(b) != 1 {
		t.Fatalf("invalid not output len, got %d, expected 1", len(b))
	}
	if !strings.Contains(b[0], expected) {
		t.Errorf("log error, got %q, does not cotain expected %q", b[0], expected)
	}
}

func TestPanic(t *testing.T) {
	testlog := "testPanic.log"
	expected := "Panic test"

	oldStderr := os.Stderr
	_, stderrW, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stderr = stderrW

	defer func() {
		if r := recover(); r != nil {

			panicStr, ok := r.(string)
			if !ok {
				t.Errorf("Expected panic with error type, got %T", r)
			} else {
				if !strings.Contains(panicStr, expected) {
					t.Errorf("Expected panic message '%s', got '%s'", expected, panicStr)
				}

				stderrW.Close()
				os.Stderr = oldStderr

				file, err := os.OpenFile(testlog, os.O_RDONLY, 0666)
				if err != nil {
					t.Fatalf("log file open failed, %q", err)
				}

				var b []string
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					b = append(b, scanner.Text())
				}
				file.Close()
				os.Remove(testlog)

				if len(b) != 1 {
					t.Fatalf("invalid not output len, got %d, expected 1", len(b))
				}
				if !strings.Contains(b[0], expected) {
					t.Errorf("log error, got %q, does not cotain expected %q", b[0], expected)
				}
			}
		} else {
			t.Error("Expected panic, but no panic occurred")
		}
	}()

	// Test code begin
	gologger.InitFileLoger(false, true, testlog)
	gologger.Panic("Panic test\n")
	// Test code end
}

func TestPanicf(t *testing.T) {
	testlog := "testPanicf.log"
	expected := "Panicf test"

	oldStderr := os.Stderr
	_, stderrW, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stderr = stderrW

	defer func() {
		if r := recover(); r != nil {

			panicStr, ok := r.(string)
			if !ok {
				t.Errorf("Expected panic with error type, got %T", r)
			} else {
				if !strings.Contains(panicStr, expected) {
					t.Errorf("Expected panic message '%s', got '%s'", expected, panicStr)
				}

				stderrW.Close()
				os.Stderr = oldStderr

				file, err := os.OpenFile(testlog, os.O_RDONLY, 0666)
				if err != nil {
					t.Fatalf("log file open failed, %q", err)
				}

				var b []string
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					b = append(b, scanner.Text())
				}
				file.Close()
				os.Remove(testlog)

				if len(b) != 1 {
					t.Fatalf("invalid not output len, got %d, expected 1", len(b))
				}
				if !strings.Contains(b[0], expected) {
					t.Errorf("log error, got %q, does not cotain expected %q", b[0], expected)
				}
			}
		} else {
			t.Error("Expected panic, but no panic occurred")
		}
	}()

	// Test code begin
	gologger.InitFileLoger(false, true, testlog)
	gologger.Panicf("%s\n", "Panicf test")
	// Test code end
}
func TestPanicln(t *testing.T) {
	testlog := "testPanicln.log"
	expected := "Panicln test"

	oldStderr := os.Stderr
	_, stderrW, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stderr = stderrW

	defer func() {
		if r := recover(); r != nil {

			panicStr, ok := r.(string)
			if !ok {
				t.Errorf("Expected panic with error type, got %T", r)
			} else {
				if !strings.Contains(panicStr, expected) {
					t.Errorf("Expected panic message '%s', got '%s'", expected, panicStr)
				}

				stderrW.Close()
				os.Stderr = oldStderr

				file, err := os.OpenFile(testlog, os.O_RDONLY, 0666)
				if err != nil {
					t.Fatalf("log file open failed, %q", err)
				}

				var b []string
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					b = append(b, scanner.Text())
				}
				file.Close()
				os.Remove(testlog)

				if len(b) != 1 {
					t.Fatalf("invalid not output len, got %d, expected 1", len(b))
				}
				if !strings.Contains(b[0], expected) {
					t.Errorf("log error, got %q, does not cotain expected %q", b[0], expected)
				}
			}
		} else {
			t.Error("Expected panic, but no panic occurred")
		}
	}()

	// Test code begin
	gologger.InitFileLoger(false, true, testlog)
	gologger.Panicln("Panicln test")
	// Test code end
}
