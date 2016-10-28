package update

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestReadWatchList(t *testing.T) {
	const fname = "TestReadWatchList"
	t.Log(fname, "START")
	watchList, err := readWatchList("testdata/TestReadWatchList")
	if err != nil {
		t.Fatal(fname, err)
	}
	log.Println(fname, watchList)
	t.Log(fname, "END")
}

func TestToMoviesArray(t *testing.T) {
	const fname = "TestToMoviesArray"
	t.Log(fname, "START")
	moviesArray, err := toMoviesArray(readWatchList("testdata/TestReadWatchList"))
	if err != nil {
		t.Fatal(fname, err)
	}
	log.Println(fname, moviesArray)
	t.Log(fname, "END")
}

func TestReadMoviesPersistenceJSON(t *testing.T) {
	const fname = "TestReadMoviesPersistenceJSON"
	t.Log(fname, "START")
	movies, err := readMoviesPersistenceJSON("testdata/TestReadMoviesPersistenceJSON/movies-persistence.json")
	if err != nil {
		t.Fatal(fname, err)
	}
	log.Println(fname, movies)

	t.Log(fname, "END")
}

func TestMerge(t *testing.T) {
	const fname = "TestMerge"
	t.Log(fname, "START")
	moviesArray, lerr := toMoviesArray(readWatchList("testdata/TestMerge"))
	if lerr != nil {
		t.Fatal(fname, lerr)
	}
	movies, jerr := readMoviesPersistenceJSON("testdata/TestMerge/movies-persistence.json")
	if jerr != nil {
		t.Fatal(fname, jerr)
	}
	err := merge(moviesArray, movies, "testdata/TestMerge/movies-persistence.json")
	if err != nil {
		t.Fatal(fname, err)
	}

	file, _ := os.Open("testdata/TestMerge/movies-persistence.json")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
	}
	defer file.Close()

	t.Log(fname, "END")
}
