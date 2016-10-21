package client

// アプリの終了コード
const (
	ExitCodeOK int = iota
	ExitCodeArgsError
	ExitCodeLogSetupError
	ExitCodeConfigError
	ExitCodeMovieError
	ExitCodeError
	ExitCodePanic
)
