package update

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	moviess2p "github.com/sky0621/work-go-movies/grpcs2p"
)

const watchListFile = "watch.list"

// IsRunning ... 処理中かどうか
var IsRunning bool

// Run ...
func Run(watchListDir string, storageAddr string) {
	const fname = "<>[Run]"
	log.Println(fname, "START")

	IsRunning = true
	defer func() {
		IsRunning = false
	}()

	// 今回追加分を読み込む。無いなら、その時点で終了。
	watchList, wlErr := toMoviesArray(readWatchList(watchListDir))
	if wlErr != nil {
		return
	}

	// 既存のWeb公開用のJSONを読み込む。
	moviesPersistence, jsErr := readMoviesPersistenceJSON(storageAddr)
	if jsErr != nil {
		return
	}

	// 既存のWeb公開用のJSONに対して今回追加分を足しこみ、保存
	err := merge(watchList, moviesPersistence, storageAddr)
	if err != nil {
		return
	}

	// 処理が正常に終わったら今回分の watch.list は削除
	err = os.Remove(filepath.Join(watchListDir, watchListFile))
	if err != nil {
		log.Println(fname, err)
		return
	}

	log.Println(fname, "END")
}

func readWatchList(watchListDir string) (*os.File, error) {
	const fname = "readWatchList"
	log.Println(fname, "START")
	file, err := os.Open(filepath.Join(watchListDir, watchListFile))
	if err != nil {
		log.Println(fname, err)
		return nil, err
	}
	log.Println(fname, "END")
	return file, nil
}

func toMoviesArray(watchList *os.File, err error) ([]*moviess2p.Movie, error) {
	const fname = "toMoviesArray"
	log.Println(fname, "START")
	if err != nil {
		return nil, err
	}
	defer watchList.Close()

	var moviesArray []*moviess2p.Movie
	scanner := bufio.NewScanner(watchList)
	for scanner.Scan() {
		list := scanner.Text()
		lists := strings.Split(list, "\t")
		movie := &moviess2p.Movie{
			Skey:          "1000000X",
			Filename:      lists[0],
			Title:         "",
			Playtime:      "1",
			Photodatetime: lists[1]}
		moviesArray = append(moviesArray, movie)
	}
	serr := scanner.Err()
	if serr != nil {
		log.Println(fname, serr)
		return nil, serr
	}
	log.Println(fname, "END")
	return moviesArray, nil
}

func readMoviesPersistenceJSON(storageAddr string) (*moviess2p.Movies, error) {
	const fname = "readMoviesPersistenceJSON"
	log.Println(fname, "START")
	file, err := ioutil.ReadFile(storageAddr)
	if err != nil {
		log.Println(fname, err)
		return nil, err
	}
	var resS2pMovies moviess2p.Movies
	err = json.Unmarshal(file, &resS2pMovies)
	if err != nil {
		log.Println(fname, err)
		return nil, err
	}
	log.Println(fname, "END")
	return &resS2pMovies, nil
}

func merge(watchList []*moviess2p.Movie, moviesPersistence *moviess2p.Movies, storageAddr string) error {
	const fname = "merge"
	log.Println(fname, "START")
	moviesPersistence.Movies = append(watchList, moviesPersistence.Movies...)
	jsonByteArray, err := json.MarshalIndent(moviesPersistence, "", "  ")
	if err != nil {
		log.Println(fname, err)
		return err
	}
	ioutil.WriteFile(storageAddr, jsonByteArray, os.ModePerm)
	log.Println(fname, "END")
	return nil
}
