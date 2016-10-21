package main

import (
	"flag"
	"os"

	c "github.com/sky0621/work-go-movies/client"
)

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

	os.Exit(c.Run(arg))
}

func parseFlag() (*c.Arg, error) {
	// TODO アプリバージョンの表示「-version」も入れる

	var addr string
	var logDir string
	var isDebug bool
	flag.StringVar(&addr, "addr", ":7010", "エンドポイントのアドレス")
	flag.StringVar(&addr, "a", ":7010", "エンドポイントのアドレス")
	flag.StringVar(&logDir, "log", ".", "ログ出力先ディレクトリ")
	flag.StringVar(&logDir, "l", ".", "ログ出力先ディレクトリ")
	flag.BoolVar(&isDebug, "debug", false, "デバッグモード")
	flag.BoolVar(&isDebug, "d", false, "デバッグモード")
	flag.Parse()

	// TODO NewArg内でバリデーション実装後、errが返る可能性が発生。
	arg, err := c.NewArg(addr, logDir, isDebug)
	if err != nil {
		return nil, err
	}

	return arg, nil
}
