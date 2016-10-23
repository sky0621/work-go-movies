package persistence

// アプリの終了コード
const (
	ExitCodeOK int = iota
	ExitCodeArgsError
	ExitCodeLogSetupError
	ExitCodeConfigError
	ExitCodeGRPCError
	ExitCodeMovieError
	ExitCodeError
	ExitCodePanic
)
