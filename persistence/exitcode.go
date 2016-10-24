package persistence

// アプリの終了コード
const (
	ExitCodeOK int = iota
	ExitCodeArgsError
	ExitCodeLogSetupError
	ExitCodeConfigError
	ExitCodeStorageError
	ExitCodeGRPCError
	ExitCodeMovieError
	ExitCodeError
	ExitCodePanic
)
