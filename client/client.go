package client

import movies "github.com/sky0621/work-go-movies"

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

	exitCode := webapiProvide(arg, movies.NewMovieServiceClient(grpcConn))
	if exitCode != ExitCodeOK {
		return exitCode
	}

	applog.debug(fname, "END")
	return ExitCodeOK
}
