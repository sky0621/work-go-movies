package client

import (
	"net/http"
	"sync"
)

var (
	propertyLock sync.RWMutex
	propertyMap  map[*http.Request]map[string]interface{} // あらゆるものをハンドラー間で受け渡しできるよう定義（例：GRPC接続クライアントオブジェクト）
)

// リクエストごとの共有マップを生成
func openPropertyMap(req *http.Request) {
	propertyLock.Lock()
	if propertyMap == nil {
		propertyMap = map[*http.Request]map[string]interface{}{}
	}
	propertyMap[req] = map[string]interface{}{}
	propertyLock.Unlock()
}

// リクエストが終わったら共有マップから削除
func closePropertyMap(req *http.Request) {
	propertyLock.Lock()
	delete(propertyMap, req)
	propertyLock.Unlock()
}

func getProperty(req *http.Request, key string) interface{} {
	propertyLock.RLock() // 書き込みはダメだけど同時読み出しはOK
	value := propertyMap[req][key]
	propertyLock.RUnlock()
	return value
}

func setProperty(req *http.Request, key string, value interface{}) {
	propertyLock.Lock()
	propertyMap[req][key] = value
	propertyLock.Unlock()
}
