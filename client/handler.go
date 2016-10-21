package client

import (
	"log"
	"net/http"
)

// /movies/ をハンドリング
func handleMovies(w http.ResponseWriter, r *http.Request) {
	// GET /movies/ ... 全動画をリターン
	// GET /movies/{id} ... id で特定した動画をリターン
	// POST /movies/ ... 動画を登録
	// PUT /movies/{id} ... id で特定した動画を更新
	// DELETE /movies/{id} ... id で特定した動画を削除（※全動画削除はさすがに用意しない）
	log.Println("<handler.go>[handleMovies]START")
}
