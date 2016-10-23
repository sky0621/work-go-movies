package main

import (
	"flag"
	"os"

	s "github.com/sky0621/work-go-movies/server"
)

// 【server概要】
// エンドユーザ（※ビュープログラム含む）に対してWebAPIを提供する。
// [client]とのGRPC通信[grpcc2s/movie_c2s.go使用]によって要求された動画リソース要素のCRUDに応じて
// [persistence]とGRPC通信[grpcs2p/movie_s2p.go使用]し、必要な情報のCRUDを行う。
// ■[persistence]とのGRPC通信で得た動画リソース要素を[client]が欲する形式に変換（逆も）する。
// ■[server]としてやるべきチェック処理等。
// ＜関心事＞
// ・どの[client]と接続してクライアントへの情報提供を行うか
// ・どの[persistence]と接続して動画リソース要素の永続化を行うか
// ＜備考＞
// ・現状、ただの土管になっている感覚（[server]としての付加要素は要検討）

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

	var grpc2cPort string
	var grpc2pPort string
	var logDir string
	var isDebug bool
	flag.StringVar(&grpc2cPort, "grpc2c", ":7110", "GRPC(対 クライアント)接続先ポート")
	flag.StringVar(&grpc2cPort, "gc", ":7110", "GRPC(対 クライアント)接続先ポート")
	flag.StringVar(&grpc2pPort, "grpc2p", ":7120", "GRPC(対 永続化層)接続先ポート")
	flag.StringVar(&grpc2pPort, "gp", ":7120", "GRPC(対 永続化層)接続先ポート")
	flag.StringVar(&logDir, "log", ".", "ログ出力先ディレクトリ")
	flag.StringVar(&logDir, "l", ".", "ログ出力先ディレクトリ")
	flag.BoolVar(&isDebug, "debug", false, "デバッグモード")
	flag.BoolVar(&isDebug, "d", false, "デバッグモード")
	flag.Parse()

	// TODO NewArg内でバリデーション実装後、errが返る可能性が発生。
	arg, err := s.NewArg(grpc2cPort, grpc2pPort, logDir, isDebug)
	if err != nil {
		return nil, err
	}

	return arg, nil
}
