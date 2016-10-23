package main

import (
	"flag"
	"os"

	s "github.com/sky0621/work-go-movies/server"
)

// server
func main() {
	arg, err := parseFlag()
	if err != nil {
		os.Exit(s.ExitCodeArgsError)
	}

	logfile, err := s.SetupLog(arg.LogDir)
	if err != nil {
		os.Exit(s.ExitCodeLogSetupError)
	}
	defer logfile.Close()

	os.Exit(s.Exec(arg))
}

func parseFlag() (*s.Arg, error) {
	// TODO アプリバージョンの表示「-version」も入れる

	var grpcport string
	var logDir string
	var isDebug bool
	flag.StringVar(&grpcport, "grpc", ":7110", "GRPC接続先ポート")
	flag.StringVar(&grpcport, "g", ":7110", "GRPC接続先ポート")
	flag.StringVar(&logDir, "log", ".", "ログ出力先ディレクトリ")
	flag.StringVar(&logDir, "l", ".", "ログ出力先ディレクトリ")
	flag.BoolVar(&isDebug, "debug", false, "デバッグモード")
	flag.BoolVar(&isDebug, "d", false, "デバッグモード")
	flag.Parse()

	// TODO NewArg内でバリデーション実装後、errが返る可能性が発生。
	arg, err := s.NewArg(grpcport, logDir, isDebug)
	if err != nil {
		return nil, err
	}

	return arg, nil
}
