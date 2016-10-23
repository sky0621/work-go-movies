package main

import (
	"flag"
	"os"

	c "github.com/sky0621/work-go-movies/client"
)

// 【client概要】
// エンドユーザ（※ビュープログラム含む）に対してWebAPIを提供する。
// エンドユーザからの要求に応じて[server]とGRPC通信[grpcc2s/movie_c2s.go使用]して動画リソース要素のCRUD実行。
// ＜関心事＞
// ・エンドユーザに対してどのような形式（デフォルトJSON）で動画リソース要素を提供するか
// ・WebAPIとしてのセキュリティ要件（APIキーのチェックなど）
// ＜備考＞
// ・実用に耐えるビューの加工（JSフレームワーク使用を予定）は別プロジェクトが担う。
// ・認証/認可関連は別プロジェクトが担う。

// TODO シグナル受信してアプリ停止するロジックを入れる

// client
func main() {
	arg, err := parseFlag()
	if err != nil {
		os.Exit(c.ExitCodeArgsError)
	}

	logfile, err := c.SetupLog(arg.LogDir)
	if err != nil {
		os.Exit(c.ExitCodeLogSetupError)
	}
	defer logfile.Close()

	os.Exit(c.Exec(arg))
}

func parseFlag() (*c.Arg, error) {
	// TODO アプリバージョンの表示「-version」も入れる

	var grpcport string
	var addr string
	var logDir string
	var isDebug bool
	flag.StringVar(&grpcport, "grpc", ":7110", "GRPC接続先ポート")
	flag.StringVar(&grpcport, "g", ":7110", "GRPC接続先ポート")
	flag.StringVar(&addr, "addr", ":7010", "WebAPIエンドポイントのアドレス")
	flag.StringVar(&addr, "a", ":7010", "WebAPIエンドポイントのアドレス")
	flag.StringVar(&logDir, "log", ".", "ログ出力先ディレクトリ")
	flag.StringVar(&logDir, "l", ".", "ログ出力先ディレクトリ")
	flag.BoolVar(&isDebug, "debug", false, "デバッグモード")
	flag.BoolVar(&isDebug, "d", false, "デバッグモード")
	flag.Parse()

	// TODO NewArg内でバリデーション実装後、errが返る可能性が発生。
	arg, err := c.NewArg(grpcport, addr, logDir, isDebug)
	if err != nil {
		return nil, err
	}

	return arg, nil
}
