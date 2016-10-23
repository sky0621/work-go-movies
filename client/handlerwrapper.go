package client

import (
	"net/http"

	moviesc2s "github.com/sky0621/work-go-movies/grpcc2s"
)

const (
	queryparamAPIKey      = "apikey"
	propertyKeyGRPCClient = "grpcClient"
)

// 今は不要だけど、APIキーのチェック用に定義
func needAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO HTTPヘッダーからの取得にした方がいいか。
		apiKey := r.URL.Query().Get(queryparamAPIKey)
		applog.debugf("needAPIKey", "APIキー：%s", apiKey)
		if !isValidAPIKey(apiKey) {
			applog.info("needAPIKey", "APIキーが不正:"+apiKey)
			// TODO 動確重視のためコメントアウト中。先のHTTPヘッダー化と合わせて対応。
			// respondErr(w, r, http.StatusUnauthorized, "APIキーが不正です")
		}
		fn(w, r)
	}
}

// GRPC接続クライアントの保持
func withGRPCConnect(grpcClient moviesc2s.MovieC2SServiceClient, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setProperty(r, propertyKeyGRPCClient, grpcClient)
		fn(w, r)
	}
}

func withShareProperty(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		openPropertyMap(r)
		defer closePropertyMap(r)
		fn(w, r)
	}
}
