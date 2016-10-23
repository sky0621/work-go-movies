package server

import moviess2p "github.com/sky0621/work-go-movies/grpcs2p"

var applog *logger

// Exec ... 永続化層へのGRPC接続、クライアントへのGRPCリッスンを行う
func Exec(arg *Arg) int {
	const fname = "Exec"
	applog = &logger{isDebugEnable: arg.IsDebug}
	applog.debug(fname, "プログラム引数", *arg)
	applog.debug(fname, "START")

	// 対 永続化層
	applog.debug(fname, "永続化層とのコネクト開始")
	grpcConn, err := grpcConnect(arg)
	if err != nil {
		return ExitCodeGRPCError
	}
	applog.debug(fname, "永続化層とのコネクト完了")
	defer grpcConn.Close()

	s2pClient := moviess2p.NewMovieS2PServiceClient(grpcConn)

	// 対 クライアント
	applog.debug(fname, "クライアントとのコネクト開始")
	err = grpcListen(arg, s2pClient)
	if err != nil {
		return ExitCodeGRPCError
	}
	applog.debug(fname, "クライアントとのコネクト完了")

	applog.debug(fname, "END")
	return ExitCodeOK
}
