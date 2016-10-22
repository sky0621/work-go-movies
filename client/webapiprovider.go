package client

import (
	"net/http"
	"time"

	movies "github.com/sky0621/work-go-movies"
	"github.com/stretchr/graceful"
)

// Run ... 実クライアントからのCRUD要求に応じて、server にリクエストし、結果を編集してクライアントに返す。
// 返却形式はJSON。ユーザビューは別プロジェクトとして作る。
func webapiProvide(arg *Arg, grpcClient movies.MovieServiceClient) int {
	const fname = "webapiProvide"
	applog.debug(fname, "START")

	mux := http.NewServeMux()
	// 下記、デコレータ―パターンで組み合わせる
	// １．ハンドラー間での共有データ受け渡し用ハンドラー
	// 2.GRPC接続クライアント保持用ハンドラー
	// 3.APIキーのチェック用ハンドラー（※現状は不要だけど念のため）
	// ※ストレージとの接続はserverに持たせるので、ここでは不要
	mux.HandleFunc("/movies/", withShareProperty(withGRPCConnect(grpcClient, needAPIKey(handleMovies)))) // TODO リクエストのバリデーションもハンドラー追加かな

	applog.infof(fname, "Webサーバーを開始します。 接続先：%s", arg.Addr)
	graceful.Run(arg.Addr, 1*time.Second, mux)
	applog.info(fname, "Webサーバーを停止します。")

	applog.debug(fname, "END")
	return ExitCodeOK
}
