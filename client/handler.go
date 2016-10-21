package client

import (
	"io"
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
	default:
		io.WriteString(w, "GETのみ受け付けます")
	}
	applog.debug(fname, "END")
}

func handleMoviesGET(w http.ResponseWriter, r *http.Request) {
	applog.debug("handleMoviesGET", "START")
}
