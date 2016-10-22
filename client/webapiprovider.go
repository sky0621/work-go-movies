package client

import (
	"net/http"
	"time"

	"github.com/stretchr/graceful"
)

// Run ... 実クライアントからのCRUD要求に応じて、server にリクエストし、結果を編集してクライアントに返す。
// 返却形式はJSON。ユーザビューは別プロジェクトとして作る。
func webapiProvide(arg *Arg) int {
	const fname = "webapiProvide"
	applog.debug(fname, "START")

	mux := http.NewServeMux()
	// TODO ハンドラー間での値の受け渡し用ハンドラー、及び、APIキーのチェック用ハンドラーをデコレータ―パターンで組み合わせる(※ストレージとの接続はserverに持たせるので、ここでは不要)
	mux.HandleFunc("/movies/", handleMovies)

	applog.infof(fname, "Webサーバーを開始します。 接続先：%s", arg.Addr)
	graceful.Run(arg.Addr, 1*time.Second, mux)
	applog.info(fname, "Webサーバーを停止します。")

	applog.debug(fname, "END")
	return ExitCodeOK
}
