package client

import moviesc2s "github.com/sky0621/work-go-movies/grpcc2s"

var applog *logger

// Exec ... GRPC接続やWebAPIサーバ起動を行う
func Exec(arg *Arg) int {
	const fname = "Exec"
	applog = &logger{isDebugEnable: arg.IsDebug}
	applog.debug(fname, "プログラム引数", *arg)
	applog.debug(fname, "START")

	grpcConn, err := grpcConnect(arg)
	if err != nil {
		return ExitCodeGRPCError
	}
	defer grpcConn.Close()

	client := moviesc2s.NewMovieC2SServiceClient(grpcConn)

	exitCode := webapiProvide(arg, client)
	if exitCode != ExitCodeOK {
		return exitCode
	}

	applog.debug(fname, "END")
	return ExitCodeOK
}
