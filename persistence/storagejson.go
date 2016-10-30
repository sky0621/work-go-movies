package persistence

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	moviess2p "github.com/sky0621/work-go-movies/grpcs2p"
)

// StorageJSON ...
type StorageJSON struct {
	JSONPath string
	Logger   logger
}

// OpenStorage ... ファイルという性質から（DBのように）コネクションオープンして保持する必要はない。
// なので、ここでは以降の処理のため、ファイルが存在するか（なければ新規に作っておく）のチェックを行う。
func (s *StorageJSON) OpenStorage() error {
	const fname = "OpenStorage"
	s.Logger.debug(fname, "START")
	_, err := os.Stat(s.JSONPath)
	if err == nil {
		s.Logger.debug(fname, "END")
		return nil
	}
	file, err := os.Create(s.JSONPath)
	if err != nil {
		s.Logger.error(fname, err)
		return err
	}
	defer file.Close()

	s.Logger.debug(fname, "END")
	return nil
}

// CloseStorage ... ファイルという性質から（DBのように）コネクションクローズする必要はない。
func (s *StorageJSON) CloseStorage() error {
	const fname = "CloseStorage"
	s.Logger.debug(fname, "START")
	s.Logger.debug(fname, "END")
	return nil
}

// Create ...
func (s *StorageJSON) Create(param *CrudParam) (*CrudResult, error) {
	return nil, errors.New("FIXME")
}

// Read ...
func (s *StorageJSON) Read(cond *CrudCondition) (*CrudResults, error) {
	// FIXME CrudConditionでの絞り込み
	const fname = "Read"
	s.Logger.debug(fname, "START")
	file, err := ioutil.ReadFile(s.JSONPath)
	if err != nil {
		s.Logger.error(fname, err)
		return nil, err
	}
	var resS2pMovies moviess2p.Movies
	err = json.Unmarshal(file, &resS2pMovies)
	if err != nil {
		s.Logger.error(fname, err)
		return nil, err
	}
	s.Logger.debug(fname, "END")
	return &CrudResults{BindObj: resS2pMovies}, nil
}

// Update ...
func (s *StorageJSON) Update(cond *CrudCondition) (*CrudResults, error) {
	return nil, errors.New("FIXME")
}

// Delete ...
func (s *StorageJSON) Delete(cond *CrudCondition) (*CrudResults, error) {
	return nil, errors.New("FIXME")
}
