package server

// Arg ...
// TODO バリデーション
type Arg struct {
	Grpc2cPort string
	Grpc2pPort string
	LogDir     string
	IsDebug    bool
}

// NewArg ...
func NewArg(
	grpc2cPort string,
	grpc2pPort string,
	logDir string,
	isDebug bool) (*Arg, error) {
	// TODO バリデーション
	// 型変換前後でそれぞれバリデーションする用途を考慮し、構造体 Arg のオブジェクト生成関数を用意
	return &Arg{
		Grpc2cPort: grpc2cPort,
		Grpc2pPort: grpc2pPort,
		LogDir:     logDir,
		IsDebug:    isDebug,
	}, nil
}
