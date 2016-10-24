package persistence

import (
	"encoding/json"
	"io/ioutil"
	"os"

	moviess2p "github.com/sky0621/work-go-movies/grpcs2p"
)

const (
	storageName    = "movies-persistence"
	collectionName = "movies"
)

// Storage ... ストレージ操作を抽象化する構造体
type storage struct {
	// [MEMO]MongoDB考えたけどうまくデータ取得できず断念。とりあえずJSONのままファイルへ。
	// conn *mgo.Session
	storageAddr string
}

// [MEMO]結局、接続情報をflagから（かつmain()で）取得する方式では抽象化できない。。。やっぱり環境変数からの取得にすべきか。
func (s *storage) open(storageAddr string) error {
	const fname = "open"
	applog.debug(fname, "START")
	s.storageAddr = storageAddr // [MEMO]構造体生成時に渡せるけど。。。
	// var err error
	// [MEMO]MongoDB考えたけどうまくデータ取得できず断念。とりあえずJSONのままファイルへ。
	// s.conn, err = mgo.Dial(storageAddr)
	_, err := os.Stat(storageAddr)
	if err == nil {
		applog.debug(fname, "END")
		return nil
	}
	if os.IsExist(err) {
		applog.error(fname, err)
		return err
	}
	file, err := os.Create(storageAddr)
	if err != nil {
		applog.error(fname, err)
		return err
	}
	defer file.Close()
	applog.debug(fname, "END")
	return nil
}

func (s *storage) close() error {
	const fname = "close"
	applog.debug(fname, "START")
	// s.conn.Close() // TODO 何かエラー出る？
	applog.debug(fname, "END")
	return nil
}

func (s *storage) read() (*moviess2p.Movies, error) {
	const fname = "read"
	applog.debug(fname, "START")
	file, err := ioutil.ReadFile(s.storageAddr)
	if err != nil {
		applog.error(fname, err)
		return nil, err
	}
	var resS2pMovies moviess2p.Movies
	err = json.Unmarshal(file, &resS2pMovies)
	if err != nil {
		applog.error(fname, err)
		return nil, err
	}
	applog.debug(fname, "END")
	return &resS2pMovies, nil
}

func (s *storage) readOne(skey string) (*moviess2p.Movie, error) {
	const fname = "readOne"
	applog.debug(fname, "START")
	movies, err := s.read()
	if err != nil {
		applog.error(fname, err)
		return nil, err
	}
	for _, movie := range movies.Movies {
		if movie.Skey == skey {
			applog.debug(fname, "END")
			return movie, nil
		}
	}
	applog.debug(fname, "END(No Hit)")
	return nil, nil
}
