package persistence

import (
	"errors"
	"os"
	"testing"
)

func TestOpenStrage(t *testing.T) {
	t.Log("対象ファイルがあれば開けるはず")
	s := StorageJSON{JSONPath: "testdata/storagejson/OpenStorage/movies-persistence.json", Logger: logger{isDebugEnable: true}}
	err := s.OpenStorage()
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenStrage_NoFile(t *testing.T) {
	t.Log("対象ファイルがないなら作られるはず")
	s := StorageJSON{JSONPath: "testdata/storagejson/OpenStorage/movies-persistence2.json", Logger: logger{isDebugEnable: true}}
	err := s.OpenStorage()
	if err != nil {
		t.Fatal(err)
	}
	_, serr := os.Stat("testdata/storagejson/OpenStorage/movies-persistence2.json")
	if serr != nil {
		t.Fatal(serr)
	}
	defer func() {
		os.Remove("testdata/storagejson/OpenStorage/movies-persistence2.json")
	}()
}

func TestOpenStrage_Error(t *testing.T) {
	t.Log("対象ファイルがあっても開けなかったらエラーになるはず")
	s := StorageJSON{JSONPath: "testdata/storagejson/OpenStorage#//#/!/movies-persistence.json", Logger: logger{isDebugEnable: true}}
	err := s.OpenStorage()
	if err == nil {
		t.Fatal(err)
	}
}

func TestCloseStrage(t *testing.T) {
	t.Log("対象ファイルの有無・状態問わず処理は正常終了するはず")
	s := StorageJSON{}
	err := s.CloseStorage()
	if err != nil {
		t.Fatal(err)
	}
}

// FIXME
// func TestCreate(t *testing.T) {
// 	s := StorageJSON{}
// 	param := &CrudParam{}
// 	res, err := s.Create(param)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if res == nil {
// 		t.Fatal(errors.New("CrudResult shouldn't be nil."))
// 	}
// }

func TestRead(t *testing.T) {
	t.Log("条件指定無しの場合、対象ファイルの中身をすべて読み込めるはず")
	s := StorageJSON{JSONPath: "testdata/storagejson/Read/movies-persistence.json"}
	cond := &CrudCondition{}
	ress, err := s.Read(cond)
	if err != nil {
		t.Fatal(err)
	}
	if ress == nil {
		t.Fatal(errors.New("CrudResults shouldn't be nil."))
	}
	if ress.BindObj.Movies == nil {
		t.Fatal(errors.New("CrudResults.Movies shouldn't be nil."))
	}
	if len(ress.BindObj.Movies) != 7 {
		t.Fatal(errors.New("CrudResults.Movies count must be 7."))
	}
}

func TestRead_NoFile(t *testing.T) {
	t.Log("対象ファイルが存在しない場合、エラーになるはず")
	s := StorageJSON{JSONPath: "testdata/storagejson/Read/movies-persistence_NO_EXISTS.json"}
	cond := &CrudCondition{}
	_, err := s.Read(cond)
	if err == nil {
		t.Fatal(err)
	}
}

func TestRead_CannotParse(t *testing.T) {
	t.Log("対象ファイルがJSONとして解釈できない場合、エラーになるはず")
	s := StorageJSON{JSONPath: "testdata/storagejson/Read/movies-persistence_CANNOT_MARSHAL.json"}
	cond := &CrudCondition{}
	_, err := s.Read(cond)
	if err == nil {
		t.Fatal(err)
	}
}

// FIXME
// func TestUpdate(t *testing.T) {
// 	s := StorageJSON{}
// 	cond := &CrudCondition{}
// 	ress, err := s.Update(cond)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if ress == nil {
// 		t.Fatal(errors.New("CrudResult array shouldn't be nil."))
// 	}
// 	if len(ress) < 1 {
// 		t.Fatal(errors.New("CrudResult array shouldn't be count 0."))
// 	}
// }

// FIXME
// func TestDelete(t *testing.T) {
// 	s := StorageJSON{}
// 	cond := &CrudCondition{}
// 	ress, err := s.Delete(cond)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if ress == nil {
// 		t.Fatal(errors.New("CrudResult array shouldn't be nil."))
// 	}
// 	if len(ress) < 1 {
// 		t.Fatal(errors.New("CrudResult array shouldn't be count 0."))
// 	}
// }
