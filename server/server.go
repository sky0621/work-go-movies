package server

var applog *logger

// Exec ... GRPCリッスンを行う
func Exec(arg *Arg) int {
	const fname = "Exec"
	applog = &logger{isDebugEnable: arg.IsDebug}
	applog.debug(fname, "プログラム引数", *arg)
	applog.debug(fname, "START")

	err := grpcListen(arg)
	if err != nil {
		return ExitCodeGRPCError
	}

	applog.debug(fname, "END")
	return ExitCodeOK
}
