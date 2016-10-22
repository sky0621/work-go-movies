package client

import (
	"errors"
	"net/http"
)

// GET /movies/ ... 全動画をリターン
// GET /movies/{id} ... id で特定した動画をリターン
// POST /movies/ ... 動画を登録
// PUT /movies/{id} ... id で特定した動画を更新
// DELETE /movies/{id} ... id で特定した動画を削除（※全動画削除はさすがに用意しない）

// /movies/ をハンドリング
func handleMovies(w http.ResponseWriter, r *http.Request) {
	const fname = "handleMovies"
	applog.debug(fname, "START")

	switch r.Method {
	case "GET":
		handleMoviesGET(w, r)
	case "POST":
		handleMoviesPOST(w, r)
	case "PUT":
		handleMoviesPUT(w, r)
	case "DELETE":
		handleMoviesDELETE(w, r)
	default:
		respondHTTPErr(w, r, http.StatusNotFound)
	}
	applog.debug(fname, "END")
}

func handleMoviesGET(w http.ResponseWriter, r *http.Request) {
	const fname = "handleMoviesGET"
	applog.debug(fname, "START")
	p := NewPath(r.URL.Path)
	if p.HasID() {
		applog.debugf(fname, "ID: %s", p.ID)
	} else {
		applog.debugf(fname, "Path: %s", p.Path)
	}
	applog.debug(fname, "END")
}

func handleMoviesPOST(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("未実装です"))
}

func handleMoviesPUT(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("未実装です"))
}

func handleMoviesDELETE(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("未実装です"))
}
