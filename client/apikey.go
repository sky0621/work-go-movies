package client

func isValidAPIKey(key string) bool {
	// TODO 設定ファイルないしストレージにて適切なキーを管理
	return key == "apikey7777777"
}
