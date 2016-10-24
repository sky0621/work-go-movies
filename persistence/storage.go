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

func (s *storage) read() ([]*moviess2p.Movie, error) {
	const fname = "read"
	applog.debug(fname, "START")
	file, err := ioutil.ReadFile(s.storageAddr)
	if err != nil {
		applog.error(fname, err)
		return nil, err
	}
	// var movies []*moviess2p.Movie
	// var movies interface{}
	var movies map[string][]map[string]interface{}
	err = json.Unmarshal(file, &movies) // GRPC関連のデータ構造へのアンマーシャルにハマったので、ひとまずマップに詰める。
	if err != nil {
		applog.error(fname, err)
		return nil, err
	}
	// applog.debug(fname, "=======================================================================")
	// applog.debugf(fname, "%+v", movies)
	// applog.debug(fname, "=======================================================================")
	var resS2pMovies []*moviess2p.Movie
	for _, movie := range movies["movies"] {
		var resS2pMovie moviess2p.Movie
		resS2pMovie.Skey = movie["skey"].(string)
		resS2pMovie.Filename = movie["filename"].(string)
		resS2pMovie.Title = movie["title"].(string)
		resS2pMovie.Playtime = movie["playtime"].(string)
		resS2pMovie.Photodatetime = movie["photodatetime"].(string)
		resS2pMovies = append(resS2pMovies, &resS2pMovie)
	}
	// applog.debug(fname, "=======================================================================")
	// applog.debugf(fname, "%+v", resS2pMovies)
	// applog.debug(fname, "=======================================================================")
	applog.debug(fname, "END")
	// return movies, nil
	return resS2pMovies, nil
}

func (s *storage) readOne(skey string) (*moviess2p.Movie, error) {
	const fname = "readOne"
	applog.debug(fname, "START")
	movies, err := s.read()
	if err != nil {
		applog.error(fname, err)
		return nil, err
	}
	for _, movie := range movies {
		if movie.Skey == skey {
			applog.debug(fname, "END")
			return movie, nil
		}
	}
	applog.debug(fname, "END(No Hit)")
	return nil, nil
}
