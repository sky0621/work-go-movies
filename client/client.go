package client

var applog *logger

// Exec ... GRPC接続やWebAPIサーバ起動を行う
func Exec(arg *Arg) int {
	const fname = "Exec"
	applog = &logger{isDebugEnable: arg.IsDebug}
	applog.debug(fname, "プログラム引数", *arg)
	applog.debug(fname, "START")

	exitCode := grpcConnect(arg)
	if exitCode != ExitCodeOK {
		return exitCode
	}

	exitCode = webapiProvide(arg)
	if exitCode != ExitCodeOK {
		return exitCode
	}

	applog.debug(fname, "END")
	return exitCode
}
