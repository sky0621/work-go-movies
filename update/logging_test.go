package update

import "testing"

func TestSetupLog(t *testing.T) {
	t.Log("指定のパスに「movies-update.log」が存在する場合、ロガーの設定が正常終了するはず")
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
	t.Log("指定のパスに「movies-update.log」が存在しない場合、新規に作成してロガーの設定が正常終了するはず")
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
