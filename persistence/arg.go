package persistence

// Arg ...
// TODO バリデーション
type Arg struct {
	StorageAddr string
	Grpc2sPort  string
	Addr        string
	LogDir      string
	IsDebug     bool
}

// NewArg ...
// TODO 永続化ストレージとの接続情報など後で追加！
func NewArg(
	storageAddr string,
	grpc2sPort string,
	logDir string,
	isDebug bool) (*Arg, error) {
	// TODO バリデーション
	// 型変換前後でそれぞれバリデーションする用途を考慮し、構造体 Arg のオブジェクト生成関数を用意
	return &Arg{
		StorageAddr: storageAddr,
		Grpc2sPort:  grpc2sPort,
		LogDir:      logDir,
		IsDebug:     isDebug,
	}, nil
}
