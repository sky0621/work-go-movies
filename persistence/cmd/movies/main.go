package main

import (
	"flag"
	"os"

	p "github.com/sky0621/work-go-movies/persistence"
)

// persistence
func main() {
	arg, err := parseFlag()
	if err != nil {
		os.Exit(p.ExitCodeArgsError)
	}

	logfile, err := p.SetupLog(arg.LogDir)
	if err != nil {
		os.Exit(p.ExitCodeLogSetupError)
	}
	defer logfile.Close()

	os.Exit(p.Exec(arg))
}

func parseFlag() (*p.Arg, error) {
	// TODO アプリバージョンの表示「-version」も入れる

	var storageAddr string
	var grpc2sPort string
	var logDir string
	var isDebug bool
	flag.StringVar(&storageAddr, "storage", "movies-persistence.json", "ストレージサーバ接続先アドレス") // [MEMO]MongoDB考えたけどうまくデータ取得できず断念。とりあえずJSONのままファイルへ。
	flag.StringVar(&storageAddr, "s", "movies-persistence.json", "ストレージサーバ接続先アドレス")
	flag.StringVar(&grpc2sPort, "grpc2s", ":7120", "GRPC(対 サーバ)接続先ポート")
	flag.StringVar(&grpc2sPort, "gs", ":7120", "GRPC(対 サーバ)接続先ポート")
	flag.StringVar(&logDir, "log", ".", "ログ出力先ディレクトリ")
	flag.StringVar(&logDir, "l", ".", "ログ出力先ディレクトリ")
	flag.BoolVar(&isDebug, "debug", false, "デバッグモード")
	flag.BoolVar(&isDebug, "d", false, "デバッグモード")
	flag.Parse()

	// TODO NewArg内でバリデーション実装後、errが返る可能性が発生。
	arg, err := p.NewArg(storageAddr, grpc2sPort, logDir, isDebug)
	if err != nil {
		return nil, err
	}

	return arg, nil
}
