package persistence

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestSetupLog(t *testing.T) {
	t.Log("指定のパスに「movies-persistence.log」が存在する場合、ロガーの設定が正常終了するはず")
	logfile, err := SetupLog("testdata/logging/exists")
	defer logfile.Close()
	if err != nil {
		t.Fatal(err)
	}
	if logfile == nil {
		t.Fatal(err)
	}
}

func TestSetupLog_FileNotExists(t *testing.T) {
	t.Log("指定のパスに「movies-persistence.log」が存在しない場合、新規に作成してロガーの設定が正常終了するはず")
	logfile, err := SetupLog("testdata/logging/notexists")
	defer logfile.Close()
	if err != nil {
		t.Fatal(err)
	}
	if logfile == nil {
		t.Fatal(err)
	}
}

func TestSetupLog_DirNotExists(t *testing.T) {
	t.Log("指定のパス自体が存在しない場合、エラーになるはず")
	logfile, err := SetupLog("testdata/logging/notexistsdir")
	defer logfile.Close()
	if err == nil {
		t.Fatal(err)
	}
	if logfile != nil {
		t.Fatal(err)
	}
}

func TestDebug(t *testing.T) {
	t.Log("DEBUGレベルログ出力可の設定の場合、DEBUGレベルのログ出力ができるはず")
	logfile, _ := SetupLog("testdata/logging/debug")
	defer os.Remove("testdata/logging/debug/movies-persistence.log")
	defer logfile.Close()
	logr := logger{isDebugEnable: true}
	logr.debug("FunctionA", "test")
	b, err := ioutil.ReadFile("testdata/logging/debug/movies-persistence.log")
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if !strings.Contains(s, "[DEBUG][FunctionA] [test]") {
		t.Fatal(errors.New("Excepted：「[DEBUG][FunctionA] [test]」 and Actually：「" + s + "」 are difference."))
	}
}

func TestDebug_CannotLog(t *testing.T) {
	t.Log("DEBUGレベルログ出力不可の設定の場合、DEBUGレベルのログ出力ができないはず")
	logfile, _ := SetupLog("testdata/logging/debug")
	defer os.Remove("testdata/logging/debug/movies-persistence.log")
	defer logfile.Close()
	logr := logger{isDebugEnable: false}
	logr.debug("FunctionA", "test")
	b, err := ioutil.ReadFile("testdata/logging/debug/movies-persistence.log")
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if strings.Contains(s, "[DEBUG][FunctionA] [test]") {
		t.Fatal(errors.New("Actually：「" + s + "」 are written."))
	}
}

func TestDebugf(t *testing.T) {
	t.Log("DEBUGレベルログ出力可の設定の場合、DEBUGレベルのログ出力ができるはず")
	logfile, _ := SetupLog("testdata/logging/debug")
	defer os.Remove("testdata/logging/debug/movies-persistence.log")
	defer logfile.Close()
	logr := logger{isDebugEnable: true}
	logr.debugf("FunctionA", "This is a %s", "test")
	b, err := ioutil.ReadFile("testdata/logging/debug/movies-persistence.log")
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if !strings.Contains(s, "[DEBUG][FunctionA] This is a [test]") {
		t.Fatal(errors.New("Excepted：「[DEBUG][FunctionA] [test]」 and Actually：「" + s + "」 are difference."))
	}
}

func TestDebugf_CannotLog(t *testing.T) {
	t.Log("DEBUGレベルログ出力不可の設定の場合、DEBUGレベルのログ出力ができないはず")
	logfile, _ := SetupLog("testdata/logging/debug")
	defer os.Remove("testdata/logging/debug/movies-persistence.log")
	defer logfile.Close()
	logr := logger{isDebugEnable: false}
	logr.debugf("FunctionA", "This is a %s", "test")
	b, err := ioutil.ReadFile("testdata/logging/debug/movies-persistence.log")
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if strings.Contains(s, "[DEBUG][FunctionA] This is a [test]") {
		t.Fatal(errors.New("Actually：「" + s + "」 are written."))
	}
}

func TestInfo(t *testing.T) {
	t.Log("DEBUGレベルログ出力可否設定に関わらず、INFOレベルのログ出力ができるはず")
	logfile, _ := SetupLog("testdata/logging/info")
	defer os.Remove("testdata/logging/info/movies-persistence.log")
	defer logfile.Close()
	logr := logger{isDebugEnable: false}
	logr.info("FunctionA", "test")
	b, err := ioutil.ReadFile("testdata/logging/info/movies-persistence.log")
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if !strings.Contains(s, "[INFO][FunctionA] [test]") {
		t.Fatal(errors.New("Excepted：「[INFO][FunctionA] [test]」 and Actually：「" + s + "」 are difference."))
	}
}

func TestInfof(t *testing.T) {
	t.Log("DEBUGレベルログ出力可否設定に関わらず、INFOレベルのログ出力ができるはず")
	logfile, _ := SetupLog("testdata/logging/info")
	defer os.Remove("testdata/logging/info/movies-persistence.log")
	defer logfile.Close()
	logr := logger{isDebugEnable: false}
	logr.infof("FunctionA", "This is a %s", "test")
	b, err := ioutil.ReadFile("testdata/logging/info/movies-persistence.log")
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if !strings.Contains(s, "[INFO][FunctionA] This is a [test]") {
		t.Fatal(errors.New("Excepted：「[INFO][FunctionA] [test]」 and Actually：「" + s + "」 are difference."))
	}
}

func TestError(t *testing.T) {
	t.Log("DEBUGレベルログ出力可否設定に関わらず、ERRORレベルのログ出力ができるはず")
	logfile, _ := SetupLog("testdata/logging/error")
	defer os.Remove("testdata/logging/error/movies-persistence.log")
	defer logfile.Close()
	logr := logger{isDebugEnable: false}
	logr.error("FunctionA", "test")
	b, err := ioutil.ReadFile("testdata/logging/error/movies-persistence.log")
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if !strings.Contains(s, "[ERROR][FunctionA] [test]") {
		t.Fatal(errors.New("Excepted：「[ERROR][FunctionA] [test]」 and Actually：「" + s + "」 are difference."))
	}
}

func TestErrorf(t *testing.T) {
	t.Log("DEBUGレベルログ出力可否設定に関わらず、ERRORレベルのログ出力ができるはず")
	logfile, _ := SetupLog("testdata/logging/error")
	defer os.Remove("testdata/logging/error/movies-persistence.log")
	defer logfile.Close()
	logr := logger{isDebugEnable: false}
	logr.errorf("FunctionA", "This is a %s", "test")
	b, err := ioutil.ReadFile("testdata/logging/error/movies-persistence.log")
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if !strings.Contains(s, "[ERROR][FunctionA] This is a [test]") {
		t.Fatal(errors.New("Excepted：「[ERROR][FunctionA] [test]」 and Actually：「" + s + "」 are difference."))
	}
}
