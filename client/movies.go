package client

import (
	"log"
	"net/http"
	"time"

	"github.com/stretchr/graceful"
)

// Run ... 実クライアントからのCRUD要求に応じて、server にリクエストし、結果を編集してクライアントに返す。
// 返却形式はJSON。ユーザビューは別プロジェクトとして作る。
func Run(arg *Arg) {
	log.Println("<movies.go>[Run]START")

	mux := http.NewServeMux()
	mux.HandleFunc("/movies/", handleMovies)

	log.Println("<movies.go>[Run]Webサーバーを開始します：", arg.Addr)
	graceful.Run(arg.Addr, 1*time.Second, mux)
	log.Println("<movies.go>[Run]Webサーバーを停止します")
}
