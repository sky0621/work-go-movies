package client

import (
	"context"
	"errors"
	"net/http"

	moviesc2s "github.com/sky0621/work-go-movies/grpcc2s"
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
	case "OPTIONS":
		handleMoviesOPTIONS(w, r)
	default:
		respondHTTPErr(w, r, http.StatusNotFound)
	}
	applog.debug(fname, "END")
}

func handleMoviesGET(w http.ResponseWriter, r *http.Request) {
	const fname = "handleMoviesGET"
	applog.debug(fname, "START")

	client := getProperty(r, propertyKeyGRPCClient).(moviesc2s.MovieC2SServiceClient)

	p := NewPath(r.URL.Path)
	// TODO リファクタ後回し
	if p.HasID() {
		applog.debugf(fname, "ID: %s", p.ID)
		applog.debug(fname, "サーバとGetMovieで接続開始")
		resMovie, err := client.GetMovie(context.Background(), &moviesc2s.MovieSkey{Skey: p.ID})
		if err != nil {
			applog.error(fname, err)
			// TODO エラーハンドリングは後で検討
			respondErr(w, r, http.StatusNotFound, "")
			return
		}
		applog.debug(fname, "サーバとGetMovieで接続完了")
		applog.debug(fname, resMovie)

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		respond(w, r, http.StatusOK, resMovie)
	} else {
		applog.debugf(fname, "Path: %s", p.Path)
		applog.debug(fname, "サーバとGetMoviesで接続開始")
		resMovies, err := client.GetMovies(context.Background(), &moviesc2s.Movie{Skey: ""}) // TODO 全動画取得時のパラメータは再検討！
		if err != nil {
			applog.error(fname, err)
			// TODO エラーハンドリングは後で検討
			respondErr(w, r, http.StatusNotFound, "")
			return
		}
		applog.debug(fname, "サーバとGetMoviesで接続完了")
		applog.debug(fname, resMovies)

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		respond(w, r, http.StatusOK, resMovies)
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

func handleMoviesOPTIONS(w http.ResponseWriter, r *http.Request) {
	const fname = "handleMoviesOPTIONS"
	applog.debug(fname, "START")

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	respond(w, r, http.StatusOK, nil)

	applog.debug(fname, "END")
}
