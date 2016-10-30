package persistence

var applog *logger

// Exec ... server からのCRUD要求に応じて動画リソース項目を永続化ストレージを使って生成・更新・取得・削除する。
func Exec(arg *Arg) int {
	const fname = "Exec"
	applog = &logger{isDebugEnable: arg.IsDebug}
	applog.debug(fname, "プログラム引数", *arg)
	applog.debug(fname, "START")

	applog.debug(fname, "ストレージとのコネクト開始")
	var storage IStorager = &StorageJSON{JSONPath: arg.StorageAddr} // [MEMO]使うストレージの切り替えは設定ファイルに逃がすなどしたい
	err := storage.OpenStorage()
	if err != nil {
		return ExitCodeStorageError
	}
	defer storage.CloseStorage()
	applog.debug(fname, storage)
	applog.debug(fname, "ストレージとのコネクト終了")

	applog.debug(fname, "サーバとのコネクト開始")
	err = grpcListen(arg, storage)
	if err != nil {
		return ExitCodeGRPCError
	}
	applog.debug(fname, "サーバとのコネクト完了")

	applog.debug(fname, "END")
	return ExitCodeOK
}
