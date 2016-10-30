package persistence

import (
	moviess2p "github.com/sky0621/work-go-movies/grpcs2p"
)

// [MEMO]現状のJSONでのデータ保持だけでなくRDBやドキュメント指向DBなど切り替えやすくするよう、インタフェースを定義してみる。とはいえ、個人ツールではYAGNI感が強い。

// IStorager ... ストレージへの操作を定義（※ストレージへのコネクションの保持は実装する構造体が担う想定）
type IStorager interface {
	OpenStorage() error
	CloseStorage() error
	Create(param *CrudParam) (*CrudResult, error)
	Read(cond *CrudCondition) (*CrudResults, error)
	Update(cond *CrudCondition) (*CrudResults, error)
	Delete(cond *CrudCondition) (*CrudResults, error)
}

// CrudCondition ...
type CrudCondition struct {
	// 検索・更新系の条件として利用
}

// CrudParam ...
type CrudParam struct {
	// 作成系のパラメータとして利用
}

// CrudResult ...
type CrudResult struct {
	// 作成系の結果として利用
	BindObj moviess2p.Movie
}

// CrudResults ...
type CrudResults struct {
	// 検索・更新・削除系の結果として利用
	BindObj moviess2p.Movies
}
