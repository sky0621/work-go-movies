package client

// Arg ...
// TODO バリデーション
type Arg struct {
	Addr   string
	LogDir string
}

// NewArg ...
func NewArg(addr string, logDir string) *Arg {
	// TODO バリデーション
	// 型変換前後でそれぞれバリデーションする用途を考慮し、構造体 Arg のオブジェクト生成関数を用意
	return &Arg{Addr: addr, LogDir: logDir}
}
