package main

import (
	"flag"

	"github.com/sky0621/work-go-movies/client"
)

// client
func main() {
	arg := parseFlag()
	client.SetupLog(arg.LogDir)
	client.Run(arg)
}

func parseFlag() *client.Arg {
	// TODO アプリバージョンの表示「 -version 」も入れる

	var addr string
	var logDir string
	flag.StringVar(&addr, "addr", ":7010", "エンドポイントのアドレス")
	flag.StringVar(&logDir, "log", ".", "ログ出力先ディレクトリ")
	flag.Parse()
	return client.NewArg(addr, logDir)
}
